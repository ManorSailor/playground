import { FaissStore } from "@langchain/community/vectorstores/faiss";
import { EMBEDDINGS_MODEL, RAG_MODEL } from "./core.js";
import type { MessageFieldWithRole } from "@langchain/core/messages";

async function askLLM(conversation: MessageFieldWithRole[]) {
  return RAG_MODEL.invoke(conversation);
}

async function askLLMStream(conversation: MessageFieldWithRole[]) {
  return RAG_MODEL.stream(conversation);
}

async function getVectorStore() {
  return FaissStore.load("src/data/store", EMBEDDINGS_MODEL);
}

export { askLLM, askLLMStream, getVectorStore };
