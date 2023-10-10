# krill Python plugin SDK

Write custom context providers and suggest-postprocessors for krill in
Python. Plugins run as subprocesses on demand — no daemon, no polling.

## Install

```bash
pip install -e plugins/sdk
```

## Minimal provider

```python
# ~/.config/krill/plugins/my_provider.py
from krill_plugins import ContextProvider, register


@register
class JiraProvider(ContextProvider):
    name = "jira"

    def collect(self, repo: dict, file: str | None) -> str:
        ticket = repo.get("branch", "").upper()
        return f"Related ticket: {ticket} (from provider {self.name})"
```

Enable it in config:

```toml
[plugins]
enabled = true
providers = ["jira"]
```

## Why Python

The Go core is fast and single-binary; the plugin surface is the place
where you want ecosystem reach. Python providers can import any HTTP
client, parse XML/SQL, or call internal APIs — and the core never blocks
on them (2s deadline, failures are skipped silently).

## Layout

```
plugins/
├── sdk/                 # the krill_plugins package (pip-installable)
│   ├── pyproject.toml
│   └── krill_plugins/
│       ├── __init__.py
│       ├── base.py      # ContextProvider / SuggestPostprocessor
│       └── registry.py  # discovery + loading
└── examples/            # sample providers
```