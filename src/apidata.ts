import { apiSchema } from "./deps.ts";

export async function load(file: string) {
	const f = await import(file);
	console.log(f);
	return {};
}
