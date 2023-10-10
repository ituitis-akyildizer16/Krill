"""Base classes for krill plugins."""

from __future__ import annotations

from dataclasses import dataclass, field
from typing import Any, Optional


@dataclass
class ContextProvider:
    """Collects extra context that is appended to the prompt.

    Subclasses set ``name`` and implement ``collect``. The core calls
    ``collect`` with a 2-second deadline; exceptions are swallowed and
    the provider's output is skipped.
    """

    name: str = "unnamed"
    priority: int = 10

    def collect(self, repo: dict, file: Optional[str] = None) -> str:
        """Return extra context text, or "" for nothing."""
        return ""


@dataclass
class SuggestPostprocessor:
    """Post-processes a suggested shell command."""

    name: str = "unnamed"
    priority: int = 10

    def process(self, command: str, intent: str) -> str:
        """Return the (possibly modified) command."""
        return command


@dataclass
class PluginResult:
    """What one plugin produced for a command run."""

    provider_outputs: dict[str, str] = field(default_factory=dict)
    postprocessed: dict[str, str] = field(default_factory=dict)