import { defaultApiSchema } from "./deps.ts";
import { Schema } from "./schema.ts";

/**
 * load an apidata.js file
 *
 * @param file apidata.js file path
 * @returns apiSchema object
 */
export async function load(file: string | null): Promise<Schema> {
  if (file) {
    const f = await Deno.readTextFile(file);
    return eval(f + "; apiSchema");
  } else {
    return defaultApiSchema;
  }
}
