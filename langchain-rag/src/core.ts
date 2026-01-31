import { ChatOllama, OllamaEmbeddings } from "@langchain/ollama";
import type { MessageFieldWithRole } from "@langchain/core/messages";

const RAG_MODEL = new ChatOllama({
  model: process.env.RAG_MODEL!,
  baseUrl: process.env.OLLAMA_HOST!,
});

const JSON_TO_TEXT_MODEL = new ChatOllama({
  model: process.env.JSON_TO_TEXT_MODEL!,
  baseUrl: process.env.OLLAMA_HOST!,
  temperature: 0,
});

const EMBEDDINGS_MODEL = new OllamaEmbeddings({
  model: process.env.EMBEDDINGS_MODEL!,
  baseUrl: process.env.OLLAMA_HOST!,
});

const SYSTEM_PROMPT = {
  role: "system",
  content: `You're a helpful W&H Blown Film Lines (BFL) Diagnostic Support Agent created at INGSOL. You must rely on the sources for relevant information. DO NOT leak the sources to the user. Only you may rely on it and answer. If you do not know the answer despite having the context. Simply reply "I don't know. Contact Human Support."`,
} as const satisfies MessageFieldWithRole;


export {
  RAG_MODEL,
  JSON_TO_TEXT_MODEL,
  EMBEDDINGS_MODEL,
  SYSTEM_PROMPT,

};
