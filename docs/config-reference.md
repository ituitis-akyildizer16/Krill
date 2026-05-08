# Config reference

`~/.config/krill/config.toml` (override with `KRILL_CONFIG`). All keys
optional — defaults shown.

```toml
model = "qwen2.5-coder:7b"   # Ollama model for answers
shell = "bash"               # default target shell for suggest
theme = "dark"               # terminal theme (dark | light)
cache_ttl_seconds = 3600     # how long model responses stay cached
no_color = false             # disable ANSI colors

[context]
max_blame_lines = 400        # blame lines fed to the model
max_diff_lines = 600         # diff lines fed to the model
history_days = 14            # how far back recent history goes

[ollama]
url = "http://localhost:11434"
timeout_seconds = 120

[suggest]
max_tokens = 60              # max tokens for the command output
use_local_model = true       # always true; kept for clarity
deny_patterns = ["rm -rf /", "mkfs", ":(){"]
warn_patterns = ["rm ", "drop ", "git push --force", "shutdown"]

[plugins]
enabled = false
dir = "~/.config/krill/plugins"
providers = ["jira", "top_authors"]

[ignore]
paths = ["node_modules", ".git", "vendor", "dist"]
```

## Key notes

- `deny_patterns` are regex, case-insensitive, matched against the full
  suggested command. A match blocks the command entirely.
- `warn_patterns` only print a warning to stderr (suppressed with
  `--inline`).
- `plugins.providers` is the allowlist; only listed provider names load.
- Truncated context is always marked `... (truncated)` in the prompt, so
  the model knows it may be missing context.
<!-- draft note 6 -->
<!-- draft note 14 -->
