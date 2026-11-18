# TypeScript

Use TypeScript for source and configuration wherever the tool supports it. Use `.tsx` for JSX and `.ts` for other TypeScript files.

## Compiler settings

Enable strict compiler options, including `noUncheckedIndexedAccess`. Keep them in one `tsconfig.json` by default. Before introducing multiple `tsconfig.*.json` files, establish which tooling constraint requires the split with the maintainer.

## Type checking and formatting

Run `tsc --noEmit` using the project's selected package manager in `task validate`. A framework module may define a more specific validation workflow. Apply the [Biome baseline](../../shared/js-biome-core.md) and [configuration conventions](../../shared/biome-conventions.md).
