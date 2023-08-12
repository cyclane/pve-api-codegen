import { join } from "https://deno.land/std@0.198.0/path/mod.ts";
import { Schema } from "../../schema.ts";

import goMod from "./templates/go.mod.ts";

const wfOpts = {
  createNew: true,
};

export default async function (
  path: string,
  schema: Schema,
  module: string,
): Promise<void> {
  const te = new TextEncoder();
  await Promise.all([
    Deno.writeFile(join(path, "go.mod"), te.encode(goMod(module)), wfOpts),
  ]);
}
