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
