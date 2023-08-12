.PHONY: compile run apidata fmt fmt-check lint cache

compile:
	@deno compile src/main.ts -o pve-api-codegen

run:
	@deno run src/main.ts $(filter-out $@,$(MAKECMDGOALS))

apidata:
	@cp pve-docs/api-viewer/apidata.js src/apidata.js
	@sed -i 's/^const /export const /' src/apidata.js

fmt:
	@deno fmt src/

fmt-check:
	@deno fmt --check src/

lint:
	@deno lint src/

cache:
	@deno cache src/main.ts

%:
	@:
