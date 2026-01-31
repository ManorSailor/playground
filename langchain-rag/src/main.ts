import type {
  AIMessageChunk,
  MessageFieldWithRole,
} from "@langchain/core/messages";
import {
  askLLM,
  askLLMStream,
  askLLMToConvertJsonToText,
  getVectorStore,
} from "./llm.js";
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

async function rewriteJsonToText(ruleJson: string): Promise<string> {
  const prompt = `Convert the following diagnostic rule JSON into controlled expert text
using EXACTLY the structure below.

--- OUTPUT STRUCTURE ---

Diagnostic Rule ID: <rule_id>

This diagnostic rule applies to <machine_category> machines on the <oem> <platform> platform.
The affected machine module is <machine>, specifically <sub_module>.

Observed symptom:
<symptoms>

Applicable when the following conditions are present:
<conditions.present>

Apply only if the following conditions are absent:
<conditions.absent>

Unchanged conditions:
<conditions.unchanged>

Primary likely cause:
<likely_causes[priority=1]>

Secondary possible causes:
<likely_causes[priority=2,3]>

Verification test:
<Action: verification_test.action>
<Observation window: verification_test.observation_window>

Confirmation signal:
<verification_test.confirmation_signal>

Disconfirmation signal:
<verification_test.disconfirmation_signal>

Next decision path:
<next_decision_path>

Confidence level:
<confidence_level>

--- JSON INPUT ---
${JSON.stringify(ruleJson, null, 2)}
`;

  const text = await askLLMToConvertJsonToText(prompt);
  console.log("Text :",text.content);
  
  return text.text;
}

export {
  askUserQuery,
  rewriteJsonToText,
  type UserPrompt,
  type AIMessageResponse,
};
