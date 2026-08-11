# Prompt ideas

Questions that work well with krill's git grounding:

## Code understanding
- "why is this function slow?" `-f` the file
- "what does this module depend on?"
- "which callers would break if I changed the signature of X?"

## History / blame
- "who last touched the retry loop and why?"
- "what changed in this file this week?"
- "is this regression recent?"

## Before committing
- "summarize what I changed"
- "did I forget to handle an error path here?"

## Suggestions
- "stage and commit everything"
- "list largest files in this repo"
- "rebuild and rerun only failing tests"

## Review
- "does the staged diff have any risky lines?"

The best results come from questions that the context can answer —
the model is told to say "not in context" rather than guess.
<!-- draft note 11
