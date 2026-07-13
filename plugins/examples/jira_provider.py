"""Sample provider: surfaces the Jira ticket from the branch name."""

from krill_plugins import ContextProvider


class JiraProvider(ContextProvider):
    name = "jira"

    def collect(self, repo: dict, file: str | None = None) -> str:
        branch = repo.get("branch", "")
        ticket = branch.split("/")[-1].upper()
        if ticket and any(ch.isdigit() for ch in ticket):
            return f"Related ticket: {ticket}"
        return ""
# draft note 6
