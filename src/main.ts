import { load } from "./apidata.ts";
import { parse } from "./deps.ts";

const cliArgs = parse(Deno.args, {
	alias: {
		help: ["h"],
		file: ["f"]
	},
	boolean: [
		"help"
	],
	default: {
		file: "./apidata.js"
	},
	string: [
		"file"
	],
});

main(cliArgs);

function main(args: typeof cliArgs) {
	if (args.help) {
		console.log(`Proxmox API codegen

Generate a typed API library given a Proxmox apidata.js file.

Usage: pve-api-codegen ([options]) [language] ([destination])

Options:
  -h, --help
	       Show this message
	-f, --file
	       apidata.js file to use
				 Default: ./apidata.js

Language:
      One of: go

Destination:
		The destination directory for the generated library.
    Default: [language]/
`)
		return;
	}
	console.log(load(args.file));
}

