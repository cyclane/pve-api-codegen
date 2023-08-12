.PHONY: compile run apidata fmt fmt-check lint type-check cache

PERMISSIONS=--allow-run --allow-read --allow-write
IGNORE_CHECKS=--ignore=src/_apidata.ts

compile:
	@deno compile ${PERMISSIONS} src/main.ts -o pve-api-codegen

run:
	@deno run ${PERMISSIONS} -- src/main.ts $(filter-out $@,$(MAKECMDGOALS))

apidata:
	@cp pve-docs/api-viewer/apidata.js src/_apidata.ts
	@sed -i 's/^const /import \{ Schema \} from "\.\/schema.ts"; export const /' src/_apidata.ts
	@sed -i 's/^\]$$/] as Schema/' src/_apidata.ts

fmt:
	@deno fmt ${IGNORE_CHECKS} src/

fmt-check:
	@deno fmt ${IGNORE_CHECKS} --check src/

lint:
	@deno lint ${IGNORE_CHECKS} src/

type-check:
	@deno check --all src/main.ts

cache:
	@deno cache src/main.ts

%:
	@:
