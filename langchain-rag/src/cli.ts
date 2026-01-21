import * as readline from "node:readline/promises";

const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout,
});

async function streamToStdOut<T extends ReadableStream, A>(
  stream: T,
  onChunk: (chunk: A) => string,
  opts?: { padStartChunk?: string },
): Promise<string> {
  let allChunks: string = "";

  if (opts?.padStartChunk) {
    process.stdout.write(opts.padStartChunk);
  }

  for await (const d of stream) {
    const chunk = onChunk(d);
    process.stdout.write(chunk);
    allChunks += chunk;
  }

  return allChunks;
}

function readFromCLI(
  onInput: (query: string) => Promise<void> | void,
  onClose?: () => void,
) {
  const forceClose = () => process.exit(0);

  rl.addListener("line", onInput);
  rl.addListener("close", onClose ?? forceClose);
}

export { readFromCLI, streamToStdOut };
