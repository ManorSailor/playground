import { FaissStore } from "@langchain/community/vectorstores/faiss";
import { EMBEDDINGS_MODEL, JSON_TO_TEXT_MODEL, RAG_MODEL } from "./core.js";
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

async function askLLMToConvertJsonToText(prompt: string) {
  return JSON_TO_TEXT_MODEL.invoke(prompt);
}

export { askLLM, askLLMStream, getVectorStore, askLLMToConvertJsonToText };
