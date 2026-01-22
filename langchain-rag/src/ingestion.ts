import { PDFLoader } from "@langchain/community/document_loaders/fs/pdf";
import { RecursiveCharacterTextSplitter } from "@langchain/textsplitters";
import { FaissStore } from "@langchain/community/vectorstores/faiss";
import { EMBEDDINGS_MODEL } from "./core.js";
import { DirectoryLoader } from "@langchain/classic/document_loaders/fs/directory";

const loader = new DirectoryLoader("src/data", {
  ".pdf": (pdf) => new PDFLoader(pdf),
});

const docs = await loader.load();

const textSplitter = new RecursiveCharacterTextSplitter({
  chunkSize: 500,
  chunkOverlap: 100,
});

const splittedDocs = await textSplitter.splitDocuments(docs);

const vectorStore = await FaissStore.fromDocuments(splittedDocs, EMBEDDINGS_MODEL);

vectorStore.save("src/data/store");

export { vectorStore };
