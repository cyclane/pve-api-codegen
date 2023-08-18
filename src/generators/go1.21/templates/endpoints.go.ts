import { EnumAdder, enumResolver } from "../../../enum_resolver.ts";
import { Child, Endpoint, HTTPMethod, Schema } from "../../../schema.ts";
import { NameBuilder, newNameBuilder } from "../../../utils.ts";
import enumGo from "./enum.go.ts";

// function loadPropertiesEnums(name: string, prop: Properties, enums: Record<string, Promise<string>>, adder: Adder) {
// 	if (!prop.type || prop.type === "object") {
// 		Object.entries(prop.properties || {})
// 			.forEach(([n, p]) => loadPropertiesEnums(p, enums, adder));
// 	}
// }

async function endpoint<T extends HTTPMethod>(
  path: string,
  name: NameBuilder,
  schema: Endpoint<T>,
  enumAdder: EnumAdder,
) {
  const { fullName, baseName } = name.cleanup();
  fullName.insert(0, schema.method.toLowerCase());
  const pathParameterNames = getPathParameters(path);
  if (schema.parameters.type && schema.parameters.type !== "object") {
    throw new Error(
      `${path} expected object for parameters, got ${schema.parameters.type}`,
    );
  }
  const properties = schema.parameters.properties || {};

  // parse enums
  const parameterEnums: Record<string, Promise<string>> = {};
  Object.entries(properties).filter(([_, prop]) => {
    return Object.keys(prop).includes("enum");
  }).forEach(([name, prop]) => {
    if (prop.type !== "string" || !prop.enum) {
      throw new Error(`${path} detected enum ${name}, but type is wrong`);
    }
    parameterEnums[name] = enumAdder(
      [
        baseName.copy().add(name),
        fullName.blank().add(...fullName.all().slice(1), name), // full name without method
        fullName.copy().add(name),
        fullName.copy().add("parameter", name), // guaranteed unique
      ],
      prop.enum,
    ).then((f) => f.build());
  });
  enumAdder.close();

  // parse path parameters
  const pathParameters = await Promise.all(
    pathParameterNames.map(async (n) => {
      const prop = properties[n];
      delete properties[n];
      if (!prop) {
        throw new Error(`${path} expected path parameter ${n} to exist`);
      }
      if (Object.keys(prop).includes("children")) {
        throw new Error(
          `${path} expected path parameter ${n} to not have children`,
        );
      }
      if (!prop.type || prop.type === "object") {
        throw new Error(
          `${path} expected path parameter ${n} to not be an object`,
        );
      }
      return {
        name: name.blank().add(n).build(),
        type: (Object.keys(parameterEnums).includes(n)
          ? await parameterEnums[n]
          : toGoType(prop.type || "any")),
      };
    }),
  );

  // parse other parameters
  if (Object.keys(properties).length !== 0) {
    pathParameters.push({
      name: "parameters",
      type: fullName.build() + "Parameters",
    });
  }
  const requestParameters = await Promise.all(
    Object.entries(properties).map(async ([n, prop]) => {
      if (Object.keys(prop).includes("children")) {
        throw new Error(
          `${path} expected request parameter ${n} to not have children`,
        );
      }
      if (!prop.type || prop.type === "object") {
        throw new Error(
          `${path} expected request parameter ${n} to not be an object`,
        );
      }
      return {
        name: name.blank().add(n).build(),
        json: n,
        description: prop.description,
        type: (Number(prop.optional || 0) === 1 ? "*" : "") +
          (Object.keys(parameterEnums).includes(n)
            ? await parameterEnums[n]
            : toGoType(prop.type || "any")),
      };
    }),
  );

  // parse response properties

  return `
${
    Object.keys(properties).length !== 0
      ? `type ${fullName.build()}Parameters struct {
${
        requestParameters.map((p) =>
          `${
            p.description
              ? p.description.split("\n").map((l) => `\t// ${l}\n`).join("")
              : ""
          }\t${p.name} ${p.type} \`json:"${p.json}"\``
        ).join("\n")
      }
}`
      : ""
  }

type ${fullName.build()}Response struct {

}

func (c *Client) ${fullName.build()}(${
    pathParameters.map((p) => `${p.name} ${p.type}`)
      .join(", ")
  }) (*${fullName.build()}Response, error) {

}
`;
}

function toGoType(type: string) {
  switch (type) {
    case "string":
      return "string";
    case "boolean":
      return "bool";
    case "integer":
      return "int";
    case "number":
      return "float64";
    default:
      return "any";
  }
}

function getPathParameters(path: string) {
  const params = path.match(/\{[^\{\}]+\}/g);
  return params?.map((p) => p.slice(1, -1)) || [];
}

function handleChildren(
  parentName: NameBuilder,
  children: Child[],
  newAdder: () => EnumAdder,
): Promise<string>[] {
  return children.map((child) => {
    const name = parentName.copy().add(child.text);
    const ownEdnpoints = child.info &&
      Object.values(child.info).map(
        (e) => endpoint(child.path, name.copy(), e, newAdder()),
      );
    const childrenEndpoints = child.children &&
      handleChildren(name, child.children, newAdder);
    return [
      ...(ownEdnpoints || []),
      ...(childrenEndpoints || []),
    ];
  }).flat();
}

function makeName(s: string) {
  return s[0].toUpperCase() + s.slice(1);
}

export default async function (schema: Schema) {
  const { newAdder, resolve } = enumResolver();
  const endpointsPromise = Promise.all(handleChildren(
    newNameBuilder(makeName, ""),
    schema,
    newAdder,
  ));
  const enums = await resolve();
  const endpoints = await endpointsPromise;
  return `package pve

func StrPtr[T ~string](s T) *T {
	return &s
}

${(await Promise.all(enums.map(enumGo))).join("\n\n")}

${endpoints.join("\n")}
`;
}
