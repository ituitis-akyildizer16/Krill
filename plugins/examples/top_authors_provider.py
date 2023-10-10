"""Sample provider: adds the top contributors for the target file."""

from krill_plugins import ContextProvider


class TopAuthorsProvider(ContextProvider):
    name = "top_authors"

    def collect(self, repo: dict, file: str | None = None) -> str:
        if not file:
            return ""
        # In production this would shell out to `git shortlog -sne -- <file>`.
        return "Recent authors for this file: see git shortlog."