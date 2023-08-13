import { intersectArrays, NameBuilder, stringIntersect } from "./utils.ts";

export type Enum = {
  values: string[];
  final: Promise<NameBuilder>;
};

type InnerEnum = Enum & {
  names: NameBuilder[];
  resolve: (final: NameBuilder) => void;
};

export type EnumAdder =
  & ((names: NameBuilder[], values: string[]) => Promise<NameBuilder>)
  & {
    close: () => void;
  };

function sortStringsByLength(s1: string, s2: string) {
  return s1.length - s2.length;
}

/**
 * enumResolver helps shorten enum names by inspecting all needed enums
 * and combining any identical ones, using a shorter name.
 *
 * note this assumes all requested names are generated in the same way.
 * also assumes last request name is always unique.
 * @returns newAdder and resolve functions
 */
export function enumResolver(): {
  newAdder: () => EnumAdder;
  resolve: () => Promise<Enum[]>;
} {
  const enums: InnerEnum[] = [];
  const adderPromises: Promise<unknown>[] = [];
  return {
    /**
     * create a new adder function to add enums and block resolution
     * until adder is closed.
     * @returns adder function with .close() method.
     */
    newAdder: () => {
      const adder = async (names: NameBuilder[], values: string[]) => {
        let resolve = function (_: NameBuilder): void {
          throw new Error("final promise resolver doesn't exist");
        };
        const final = new Promise<NameBuilder>((res) => resolve = res);
        const idx = enums.push({
          names,
          values: values,
          final,
          resolve,
        }) - 1;
        return await enums[idx].final;
      };
      adder.close = function (): void {
        throw new Error("close promise resolver doesn't exist");
      };
      adderPromises.push(
        new Promise((res) => {
          adder.close = () => res(true);
        }),
      );
      return adder;
    },
    /**
     * Resolve enum names
     * @returns List of enums to generate types for
     */
    resolve: async () => {
      await Promise.all(adderPromises);
      const options: Record<string, InnerEnum[]> = {};
      const final: InnerEnum[] = [];
      enums.forEach((e) => {
        const key = e.values.toSorted().join("");
        if (options[key]) {
          options[key].push(e);
        } else {
          options[key] = [e];
        }
      });
      const takenNames = new Set<string>();
      Object.values(options).forEach((opts) => {
        const minLength = Math.min(...opts.map((o) => o.names.length));
        // default to shortest element from priority that all optionas have.
        let found = opts.map((o) => o.names[minLength - 1])
          .toSorted(
            (a, b) => sortStringsByLength(a.build(), b.build())
          )[0].copy();
        for (let i = 0; i < minLength; i++) {
          let req = opts[0].names[i];
          opts.slice(1).forEach((o) => {
            req = req.blank().add(
              ...intersectArrays(req.all(), o.names[i].all(), stringIntersect),
            );
          });
          if (!takenNames.has(req.build())) {
            found = req;
            break;
          }
        }
        // const found = opts.map((o) =>
        //   o.names.map((n) => {
        //     return {
        //       opt: o,
        //       name: n,
        //     };
        //   })
        // ).flat().toSorted((a, b) => a.name.length - b.name.length)
        //   .find(({ name }) => !taken.has(name))!;
        console.debug(
          `resolved enum ${found.build()} (${opts.length})`,
        );
        opts.forEach((o) => {
          o.resolve(found);
        });
        takenNames.add(found.build());
        final.push(opts[0]);
      });
      return final;
    },
  };
}
