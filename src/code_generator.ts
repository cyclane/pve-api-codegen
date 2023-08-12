import { Schema } from "./schema.ts";

export interface CodeGenerator {
  name: string;
  help: string;
  checkEnv: () => Promise<void>;
  codegen: (
    path: string,
    schema: Schema,
    ...args: string[]
  ) => Promise<void>;
}
