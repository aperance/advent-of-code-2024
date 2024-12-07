import { TextLineStream } from "@std/streams";

if (Deno.stdin.isTerminal()) {
  throw Error("Input data must be piped in via stdin");
}

const readable = Deno.stdin.readable.pipeThrough(
  new TextDecoderStream(),
);

export const getInput = async () => {
  let results = "";
  for await (const chunk of readable) {
    results = results + chunk;
  }
  return results;
};

export const getInputStream = () => readable.pipeThrough(new TextLineStream());
