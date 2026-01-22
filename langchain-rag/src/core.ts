import { ChatOllama, OllamaEmbeddings } from "@langchain/ollama";
import type { MessageFieldWithRole } from "@langchain/core/messages";

const RAG_MODEL = new ChatOllama({
  model: process.env.RAG_MODEL!,
  baseUrl: process.env.OLLAMA_HOST!,
});

const EMBEDDINGS_MODEL = new OllamaEmbeddings({
  model: process.env.EMBEDDINGS_MODEL!,
  baseUrl: process.env.OLLAMA_HOST!,
});

const SYSTEM_PROMPT = {
  role: "system",
  content: `You're a helpful W&H Blown Film Lines (BFL) Diagnostic Support Agent created at INGSOL.`,
} as const satisfies MessageFieldWithRole;

export { RAG_MODEL, EMBEDDINGS_MODEL, SYSTEM_PROMPT };
