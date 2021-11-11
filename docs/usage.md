# Usage guide for krill.

## Ask

Grounded answers about your repo:

```bash
krill ask "what changed in this file this week?" -f internal/git/blame.go
krill ask "explain the retry loop" --context-only   # debug the prompt
```

## Suggest

Natural language to a shell command:

```bash
krill suggest "stage and commit everything"
# git add -A && git commit -m "wip"

krill suggest --shell fish "list largest files" --inline
```

Destructive patterns are warned (stderr) or denied outright (deny
patterns in config). `--inline` suppresses the warning for keybindings.

## Review

Summarize what you are about to commit:

```bash
krill review              # staged diff (default)
