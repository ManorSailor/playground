import type {
  AIMessageChunk,
  MessageFieldWithRole,
} from "@langchain/core/messages";
import { askLLM, askLLMStream, getVectorStore } from "./llm.js";
import { readFromCLI, streamToStdOut } from "./cli.js";
import { SYSTEM_PROMPT } from "./core.js";

const sessionHistory: MessageFieldWithRole[] = [SYSTEM_PROMPT];

async function onUserInput(query: string) {
  if (!query.trim()) {
    console.log("Please provide input.");
    return;
  }

  const store = await getVectorStore();
  const context = await store.similaritySearch(query);

  const toolContext = context
    .map((c) => `Source: ${c.metadata.source}\nContext: ${c.pageContent}\n`)
    .join("\n");

  const humanMsg: MessageFieldWithRole = {
    role: "human",
    content: `${toolContext}Question: ${query}`,
  };
  sessionHistory.push(humanMsg);

  const aiResStream = await askLLMStream(sessionHistory);

  const streamedText = await streamToStdOut<typeof aiResStream, AIMessageChunk>(
    aiResStream,
    (c) => c.text,
    { padStartChunk: "\nAI> " },
  );

  const aiMsg: MessageFieldWithRole = { role: "ai", content: streamedText };
  sessionHistory.push(aiMsg);

  console.log("");
}

readFromCLI(onUserInput);

type UserPrompt = {
  message: string;
};

type AIMessageResponse = {
  answer: string;
  contexts: string[];
};

async function askUserQuery(query: string): Promise<AIMessageResponse> {
  if (!query.trim()) {
    console.log("Please provide input.");

    return {
      answer: "N/A",
      contexts: [],
    };
  }

  const store = await getVectorStore();
  const context = await store.similaritySearch(query);

  const toolContext = context.map(
    (c) => `Source: ${c.metadata.source}\nContext: ${c.pageContent}\n`,
  );

  const humanMsg: MessageFieldWithRole = {
    role: "human",
    content: `${toolContext}Question: ${query}`,
  };
  sessionHistory.push(humanMsg);

  const aiRes = await askLLM(sessionHistory);

  const aiMsg: MessageFieldWithRole = { role: "ai", content: aiRes.text };
  sessionHistory.push(aiMsg);

  return {
    answer: aiRes.text,
    contexts: toolContext,
  };
}

export { askUserQuery, type UserPrompt, type AIMessageResponse };
