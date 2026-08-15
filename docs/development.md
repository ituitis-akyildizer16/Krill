# krill development

Component conventions live here so contributors know where things belong.

## Go core

- One package per concern under `internal/`; no circular imports.
- Every `git` output is parsed, never string-compared raw.
- Context caps are enforced in `prompt`, not in `git`.
- Run `make test` (with `-race`) before pushing.

## Extension

- The extension shells out to the CLI; it must not duplicate logic.
- Commands return user-visible messages; errors surface as
  notifications.
- `npm run compile` must pass (strict TS).

## Plugins

- Providers are pure functions of (repo, file) → text.
- Never block: keep work under the 2s deadline.
- New providers ship with an example in `plugins/examples/`.

## Commits

Conventional style: `feat:`, `fix:`, `docs:`, `chore:`, `refactor:`.
One concern per commit. Reference the issue when one exists.
<!-- draft note 7 -->
<!-- draft note 15 -->
<!-- draft note 23 -->
