import { RAG_MODEL } from "./core.js";
import type { MessageFieldWithRole } from "@langchain/core/messages";

async function askLLM(conversation: MessageFieldWithRole[]) {
  return RAG_MODEL.invoke(conversation);
}

async function askLLMStream(conversation: MessageFieldWithRole[]) {
  return RAG_MODEL.stream(conversation);
}

export { askLLM, askLLMStream };
