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
