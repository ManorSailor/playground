import type {
  AIMessageChunk,
  MessageFieldWithRole,
} from "@langchain/core/messages";
import { askLLMStream } from "./llm.js";
import { readFromCLI, streamToStdOut } from "./cli.js";
import { SYSTEM_PROMPT } from "./core.js";

const sessionHistory: MessageFieldWithRole[] = [SYSTEM_PROMPT];

async function onUserInput(query: string) {
  if (!query.trim()) {
    console.log("Please provide input.");
    return;
  }

  const humanMsg: MessageFieldWithRole = { role: "human", content: query };
  sessionHistory.push(humanMsg);

  const aiResStream = await askLLMStream(sessionHistory);

  const aiText = await streamToStdOut<typeof aiResStream, AIMessageChunk>(
    aiResStream,
    (c) => c.text,
    { padStartChunk: "\nAI> " }
  );

  const aiMsg: MessageFieldWithRole = { role: "ai", content: aiText };
  sessionHistory.push(aiMsg);

  console.log("");
}

readFromCLI(onUserInput);
