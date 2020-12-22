"""Sample provider: surfaces the Jira ticket from the branch name."""

from krill_plugins import ContextProvider


class JiraProvider(ContextProvider):
    name = "jira"
