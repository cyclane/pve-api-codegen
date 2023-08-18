import { load } from "./apidata.ts";
import { CodeGenerator } from "./code_generator.ts";
import { parse } from "./deps.ts";

import go121 from "./generators/go1.21/mod.ts";

const cliArgs = parse(Deno.args, {
  alias: {
    help: ["h"],
    file: ["f"],
    output: ["o"],
    overwrite: ["O"],
  },
  boolean: [
    "help",
    "overwrite",
  ],
  default: {
    file: null,
    output: null,
    overwrite: false,
  },
  string: [
    "file",
    "output",
  ],
});

const codeGenerators: CodeGenerator[] = [
  go121,
];

async function existsNotEmpty(path: string): Promise<string | undefined> {
  try {
    const stat = await Deno.stat(path);
    if (!stat.isDirectory) {
      return "destination is not a directory";
    }
    for await (const _ of Deno.readDir(path)) {
      return "destination directory is not empty";
    }
  } catch (error) {
    if (error instanceof Deno.errors.NotFound) {
      return;
    }
    throw error;
  }
}

async function main(args: typeof cliArgs) {
  // console.debug = () => {}; // remove comment for testing
  if (args.help || args._.length === 0) {
    console.log(`Proxmox API codegen

Generate a typed API library given a Proxmox apidata.js file.

Usage:   pve-api-codegen ([options]) [language] [language_args]...
Example: pve-api-codegen -o ./go/ go1.21 github.com/cyclane/pve-go

Options:
  -h, --help
         Show this message
  -f, --file [file]
         apidata.js file to use
         Default: <bundled apidata.js file>
  -o, --output [destination]
         The destination directory for the generated library.
         Default: [language]/
  -O, --overwrite
         Overwrite existing files instead of erroring.

Available Languages:
${
      codeGenerators.map((cg) => cg.help.split("\n").map((l) => "  " + l))
        .flat().join("\n")
    }
`);
    return;
  }
  const [language, ...codeGeneratorArgs] = args._.map((v) => v.toString());
  const apiSchema = await load(args.file);
  const codeGenerator = codeGenerators.find((cg) => cg.name === language);
  if (!codeGenerator) {
    console.error(`unsupported language '${language}'`);
    return 1;
  }
  const destination = args.output || language + "/";
  const existsNotEmptyMessage = await existsNotEmpty(destination);
  if (existsNotEmptyMessage && !args.overwrite) {
    console.error(existsNotEmptyMessage);
    return 1;
  }

  await Deno.mkdir(destination, { recursive: true });
  await codeGenerator.checkEnv();
  await codeGenerator.codegen(destination, apiSchema, ...codeGeneratorArgs);
}

Deno.exit(await main(cliArgs));
