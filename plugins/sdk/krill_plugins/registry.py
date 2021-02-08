"""Discovery and loading of krill plugins from a plugins directory."""

from __future__ import annotations

import importlib.util
import sys
from pathlib import Path
from typing import Optional

from krill_plugins.base import ContextProvider, SuggestPostprocessor


def _load_module(path: Path) -> Optional[object]:
