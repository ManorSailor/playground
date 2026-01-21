import type { MessageFieldWithRole } from "@langchain/core/messages";
import { askLLM } from "./llm.js";
import { readFromCLI } from "./cli.js";
import { SYSTEM_PROMPT } from "./core.js";

const sessionHistory: MessageFieldWithRole[] = [SYSTEM_PROMPT];

async function onUserInput(query: string) {
  if (!query.trim()) {
    console.log("Please provide input.");
    return;
  }

  const humanMsg: MessageFieldWithRole = { role: "human", content: query };
  sessionHistory.push(humanMsg);

  const aiRes = await askLLM(sessionHistory);
  const aiMsg: MessageFieldWithRole = { role: "ai", content: aiRes.content };
  sessionHistory.push(aiMsg);

  console.log("")
  console.log("AI> ", aiMsg.content)
  console.log("")
}

readFromCLI(onUserInput);
