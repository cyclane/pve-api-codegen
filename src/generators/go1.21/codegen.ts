import { join } from "https://deno.land/std@0.198.0/path/mod.ts";
import { Schema } from "../../schema.ts";

import goMod from "./templates/go.mod.ts";
import loggerGo from "./templates/logger.go.ts";
import clientGo from "./templates/client.go.ts";
import endpointsGo from "./templates/endpoints.go.ts";

const wfOpts: Deno.WriteFileOptions = {
  create: true,
};

export default async function (
  path: string,
  schema: Schema,
  module: string,
): Promise<void> {
  if (!module) throw new Error("[module] is required");
  const te = new TextEncoder();
  await Promise.all([
    Deno.writeFile(join(path, "go.mod"), te.encode(goMod(module)), wfOpts),
    Deno.writeFile(join(path, "logger.go"), te.encode(loggerGo()), wfOpts),
    Deno.writeFile(
      join(path, "client.go"),
      te.encode(clientGo(module)),
      wfOpts,
    ),
    Deno.writeFile(
      join(path, "endpoints.go"),
      te.encode(await endpointsGo(schema)),
      wfOpts,
    ),
  ]);
  let command = new Deno.Command("go", {
    args: ["mod", "tidy"],
    stdout: "piped",
    stderr: "piped",
    cwd: path,
  });
  let child = command.spawn();
  if ((await child.status).code !== 0) {
    const output = await child.output();

    const td = new TextDecoder();
    const stdout = td.decode(output.stdout).trim();
    const stderr = td.decode(output.stderr).trim();
    throw new Error(
      `go mod tidy failed:\n\tSTDOUT: ${stdout}\n\tSTDERR: ${stderr}`,
    );
  }
  command = new Deno.Command("go", {
    args: ["fmt"],
    stdout: "piped",
    stderr: "piped",
    cwd: path,
  });
  child = command.spawn();
  if ((await child.status).code !== 0) {
    const output = await child.output();

    const td = new TextDecoder();
    const stdout = td.decode(output.stdout).trim();
    const stderr = td.decode(output.stderr).trim();
    throw new Error(`go fmt failed:\n\tSTDOUT: ${stdout}\n\tSTDERR: ${stderr}`);
  }
}
