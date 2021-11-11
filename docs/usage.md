# Usage guide for krill.

## Ask

Grounded answers about your repo:

```bash
krill ask "what changed in this file this week?" -f internal/git/blame.go
krill ask "explain the retry loop" --context-only   # debug the prompt
```

## Suggest

Natural language to a shell command:
