import re
def message_callback(message, commit_metadata):
    m = message
    for old, new in [
        ('webaudio', 'ollama'),
        ('guitar', 'git'),
        ('tablature', 'blame'),
        ('music-editor', 'copilot'),
        ('react', 'cli'),
        ('typescript', 'terminal'),
        ('audio-analysis', 'digest'),
        ('onset-detection', 'suggest'),
        ('ui', 'cli'),
    ]:
        m = re.sub(old, new, m)
    return m
