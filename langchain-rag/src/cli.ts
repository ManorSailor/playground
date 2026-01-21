import * as readline from "node:readline/promises";

function readFromCLI(
  onInput: (query: string) => Promise<void> | void,
  onClose?: () => void,
) {
  const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout,
  });

  const forceClose = () => process.exit(0);

  rl.addListener("line", onInput);
  rl.addListener("close", onClose ?? forceClose);
}

export { readFromCLI };
