# Libraries published to npm

Apply these choices when publishing a JavaScript or TypeScript library to npm.

| capability | status | package | source | notes |
| --- | --- | --- | --- | --- |
| npm-library-bundler | required | tsdown | https://tsdown.dev/ | Bundle the library for publication. |

## Package validation

Validate the published package layout with `publint`. When a library exposes TypeScript types, generate the declaration files and validate them with [attw](https://tsdown.dev/options/package-validation).
