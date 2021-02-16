"""Discovery and loading of krill plugins from a plugins directory."""

from __future__ import annotations

import importlib.util
import sys
from pathlib import Path
from typing import Optional

from krill_plugins.base import ContextProvider, SuggestPostprocessor


def _load_module(path: Path) -> Optional[object]:
    spec = importlib.util.spec_from_file_location(
        f"krill_plugin_{path.stem}", path
    )
    if spec is None or spec.loader is None:
        return None
    module = importlib.util.module_from_spec(spec)
    try:
        sys.modules[spec.name] = module
        spec.loader.exec_module(module)
        return module
    except Exception:
        return None


def load_plugins(
    plugins_dir: str | Path,
) -> tuple[list[ContextProvider], list[SuggestPostprocessor]]:
    """Load all plugin modules under plugins_dir.

    Returns (providers, postprocessors) ordered by priority.
    """
    base = Path(plugins_dir).expanduser()
    if not base.is_dir():
        return [], []

    providers: list[ContextProvider] = []
    postprocessors: list[SuggestPostprocessor] = []
    for path in sorted(base.rglob("*.py")):
        if path.name.startswith("_"):
            continue
        module = _load_module(path)
        if module is None:
            continue
        for attr in dir(module):
            value = getattr(module, attr)
            if isinstance(value, type):
