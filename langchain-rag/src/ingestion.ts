import { PDFLoader } from "@langchain/community/document_loaders/fs/pdf";
import { RecursiveCharacterTextSplitter } from "@langchain/textsplitters";
import { FaissStore } from "@langchain/community/vectorstores/faiss";
import { EMBEDDINGS_MODEL } from "./core.js";
import { DirectoryLoader } from "@langchain/classic/document_loaders/fs/directory";
import { BaseDocumentLoader } from "@langchain/core/document_loaders/base";
import { Document } from "@langchain/core/documents";
import fs from "fs/promises";
import { rewriteJsonToText } from "./main.js";

const DATA_DIR = "src/data/files";
const STORE_DIR = "src/data/store";

type LoaderOptions = {
  // Field to use as the main content (if not specified, stringifies the whole object)
  contentField?: string;
  // Additional metadata to include
  includeMetadata?: boolean;
  // Whether to include all object fields in metadata
  includeAllFieldsInMetadata?: boolean;
};

class JsonToTextDocument extends BaseDocumentLoader {
  constructor(
    private filePath: string,
    private options?: LoaderOptions,
  ) {
    super();
    this.filePath = filePath;
    this.options = {
      contentField: options?.contentField ?? "",
      includeMetadata: options?.includeMetadata !== false,
      includeAllFieldsInMetadata: options?.includeAllFieldsInMetadata || false,
    };
  }

  async load() {
    try {
      const fileContent = await fs.readFile(this.filePath, "utf-8");
      const jsonData = JSON.parse(fileContent);

      const dataArray = Array.isArray(jsonData) ? jsonData : [jsonData];

      // Convert each JSON object to a Document
      const documents = await Promise.all(
        dataArray.map(async (obj, index) => {
          const pageContent = await rewriteJsonToText(obj);

          // Build metadata
          const metadata = {
            source: this.filePath,
            index: index,
          };

          if (this.options?.includeAllFieldsInMetadata) {
            Object.assign(metadata, obj);
          }

          return new Document({
            pageContent,
            metadata,
          });
        }),
      );

      return documents;
    } catch (error) {
      throw new Error(`Error loading JSON file: ${error}`);
    }
  }
}

const loader = new DirectoryLoader(DATA_DIR, {
  ".pdf": (pdf) => new PDFLoader(pdf),
  ".json": (json) =>
    new JsonToTextDocument(json, { includeAllFieldsInMetadata: true }),
});

const docs = await loader.load();

const textSplitter = new RecursiveCharacterTextSplitter({
  chunkSize: 500,
  chunkOverlap: 100,
});

const splittedDocs = await textSplitter.splitDocuments(docs);

const vectorStore = await FaissStore.fromDocuments(
  splittedDocs,
  EMBEDDINGS_MODEL,
);

vectorStore.save(STORE_DIR);

export { vectorStore };
