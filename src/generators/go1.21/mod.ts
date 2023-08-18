import { CodeGenerator } from "../../code_generator.ts";
import checkEnv from "./check_env.ts";
import codegen from "./codegen.ts";

const generator: CodeGenerator = {
  name: "go1.21",
  help: `go1.21 [module]
    module: Module to use for go.mod .`,
  checkEnv,
  codegen,
};

export default generator;
