/**
 * Basic array intersection function. ( O(a1.length * a2.length) )
 * @param a1 First array.
 * @param a2 Second array.
 * @param fn Intersection function for array elements.
 * @returns Intersection of a1 and 2.
 */
export function intersectArrays<T>(
  a1: Readonly<T[]> | T[],
  a2: Readonly<T[]> | T[],
  fn: (e1: T, e2: T) => T | null,
): T[] {
  let p2 = 0;
  const out = [];
  for (const e1 of a1) {
    for (const [i2, e2] of a2.slice(p2).entries()) {
      const intT = fn(e1, e2);
      if (intT !== null) {
        out.push(intT);
        p2 += i2 + 1;
        break;
      }
    }
    if (p2 >= a2.length) break;
  }
  return out;
}

export function stringIntersect(
  s1: string,
  s2: string
) {
  if (s2.includes(s1)) return s1;
  if (s1.includes(s2)) return s2;
  return null;
}

export interface NameBuilder {
  /** Add a name to the builder */
  add: (...name: string[]) => NameBuilder;
  /** Get the name builder as an array */
  all: () => readonly string[];
  /** Build the name into a string */
  build: () => string;
  /** Copy the builder and its contents */
  copy: () => NameBuilder;
  /** Create a new builder using the same `fn` and `joinChar` */
  blank: () => NameBuilder;
  /** The make name function the builder was created with */
  fn: (name: string) => string;
}

/**
 * String array with helper functions for building a name array.
 * @param fn Make name function (used for enforcing case).
 * @param joinChar Character to join names with in .build() (used for enforcing case).
 * @returns The string array.
 */
export function newNameBuilder(
  fn: NameBuilder["fn"],
  joinChar: string,
): NameBuilder {
  const names: string[] = [];
  const builder = {
    add: (...name: string[]) => {
      names.push(
        ...name.map((n) => {
          if (n.match(/^\{[^\{\}]+\}$/)) { // match parameter
            // do not split or clean parameters just yet
            return n;
          }
          return n.split(/[^a-zA-Z0-9]/);
        }
        ).flat().filter((n) => n)
      );
      return builder;
    },
    all: () => names as readonly string[],
    build: () => names.map(fn).join(joinChar),
    copy: () => newNameBuilder(fn, joinChar).add(...names),
    blank: () => newNameBuilder(fn, joinChar),
    fn,
  };
  return builder;
}

function deplurify(word: string) {
  const defaultWhitelist = [
    "nodes",
    "groups",
    "rules",
    "resources",
    "plugins",
    "flags",
    "vnets",
    "subnets",
    "zones",
    "controllers",
    "ipams",
    "tasks",
    "users",
    "domains",
    "pools",
    "services",
    "roles",
  ];
  switch (word) {
    case "aliases":
      return "alias";
    default:
      if (!defaultWhitelist.includes(word)) {
        console.warn(`could not deplurify '${word}'`);
      }
      return word.slice(0, -1);
  }
}

/**
 * Cleanup a name.
 * @param name Name to cleanup.
 * @param method HTTP method to include in full name.
 * @returns A full name (parameters are included) and a base name (parameters are not included).
 */
export function cleanupName(name: NameBuilder, method: string): {
  fullName: NameBuilder;
  baseName: NameBuilder;
} {
  const fullName: string[] = [method.toLowerCase()];
  const baseName: string[] = [];
  const plurifyBlacklist = ["dns", "mds", "fs", "zfs"];
  let lastIsParameter = false;
  for (const n of name.all()) {
    if (n.match(/^\{[^\{\}]+\}$/)) { // match parameter
      const last = fullName[fullName.length - 1];
      if (
        !lastIsParameter && last.slice(-1)[0] === "s" &&
        !plurifyBlacklist.includes(last)
      ) {
        // shorten plural and do not include parameter
        fullName[fullName.length - 1] = deplurify(last);
        if (baseName[baseName.length - 1] === last) {
          baseName[baseName.length - 1] = deplurify(last);
        }
      } else {
        // otherwise add "By" for name to make more sense
        fullName.push("by", name.fn(n).slice(1, -1));
      }
      lastIsParameter = true;
    } else { // non-parameters
      baseName.push(n);
      fullName.push(n);
      lastIsParameter = false;
    }
  }
  return {
    fullName: name.blank().add(...fullName),
    baseName: name.blank().add(...baseName),
  };
}
