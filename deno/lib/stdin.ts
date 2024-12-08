import { TextLineStream } from "@std/streams";

if (Deno.stdin.isTerminal()) {
  throw Error("Input data must be piped in via stdin");
}

export const lineStream = Deno.stdin.readable
  .pipeThrough(new TextDecoderStream())
  .pipeThrough(new TextLineStream());
