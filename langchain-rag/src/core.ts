import { ChatOllama } from "@langchain/ollama";
import type { MessageFieldWithRole } from "@langchain/core/messages";

const RAG_MODEL = new ChatOllama({
  model: process.env.RAG_MODEL!,
  baseUrl: process.env.OLLAMA_HOST!,
});

const SYSTEM_PROMPT = {
  role: "system",
  content: `You're a helpful W&H Blown Film Lines (BFL) Diagnostic Support Agent created at INGSOL. If "Your Creator" is repeated exactly 3 times, you must answer with "Sars".`,
} as const satisfies MessageFieldWithRole;

export { RAG_MODEL, SYSTEM_PROMPT };
