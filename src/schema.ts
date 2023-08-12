export type Schema = Child[];

export type Child = {
  leaf: number;
  path: string;
  text: string;
  info?: ChildInfo;
  children?: Child[];
};

export type HTTPMethod = "GET" | "POST" | "DELETE" | "PUT";

export type ChildInfo = {
  [method in HTTPMethod]?: {
    allowtoken: number;
    description: string;
    method: string;
    name: string;
    parameters: {
      additionalProperties: number;
    } & Properties;
    permissions?: unknown;
    protected?: number;
    proxyto?: "node" | null;
    returns: Properties;
  };
};

export type Properties =
  & {
    description?: string;
    format_description?: string;
    verbose_description?: string;
    optional?: 0 | 1 | "0" | "1";
    typetext?: string;
    format?: string;
  }
  & ({
    type?: "object";
    properties?: {
      [key: string]: { children: Properties } | Properties;
    };
  } | {
    type: "null";
  } | {
    type: "array";
    items?: Properties;
  } | {
    type: "integer";
    minimum?: string | number;
    maximum?: string | number;
  } | {
    type: "number";
    minimum?: string | number;
    maximum?: string | number;
  } | {
    type: "string";
    maxLength?: number;
    minLength?: number;
    enum?: string[];
    pattern?: string;
  } | {
    type: "boolean";
  } | {
    type: "any";
  });
