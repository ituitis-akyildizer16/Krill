# krill development

Component conventions live here so contributors know where things belong.

## Go core

- One package per concern under `internal/`; no circular imports.
- Every `git` output is parsed, never string-compared raw.
- Context caps are enforced in `prompt`, not in `git`.
- Run `make test` (with `-race`) before pushing.

## Extension

- The extension shells out to the CLI; it must not duplicate logic.
