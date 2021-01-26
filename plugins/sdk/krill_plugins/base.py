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
