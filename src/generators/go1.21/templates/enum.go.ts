import { Enum } from "../../../enum_resolver.ts";

export default async function (e: Enum) {
  return `type ${(await e.final).build()} string

const (
${
    (await Promise.all(e.values.map(
      async (v) =>
        `\t${(await e.final).copy().add(v).build()} ${
          (await e.final).build()
        } = "${v}"`,
    ))).join("\n")
  }
)`;
}
