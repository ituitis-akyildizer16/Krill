# FAQ

**Why a local model instead of a cloud API?**
Privacy, latency, and cost. A 7b coding model answers most repo questions
well enough, and your diff never leaves the machine.

**Is it really fast enough?**
Context collection is ~40ms; first token ~300ms on M1 with a quantized 7b.
The model runs on your hardware, so there is no network round-trip.

**Can I use a bigger model?**
Yes — any model in Ollama: `model = "qwen2.5-coder:32b"` or a Llama 3.1
variant. Latency scales with your hardware.

**Does it work in any directory?**
Only inside a git repository (blame/diff/log need git). `krill` fails
fast with a clear message outside one.

**Does the VS Code extension need the CLI?**
Yes. The extension shells out to `krill` (`krill.binaryPath`).

**How do plugins work?**
Drop a Python file into `plugins.dir`, list its provider name in
`plugins.providers`, set `enabled = true`. Providers add context before
the prompt is assembled.

**What happens if Ollama isn't running?**
`ask`/`suggest`/`review` fail with a clear error pointing at `ollama
serve`. `status` and `version` still work.

**Is suggest safe?**
Deny patterns hard-block destructive commands (`rm -rf /`); warn patterns
flag `rm`, force-push, etc. `--inline` suppresses warnings for
keybindings — use with your own risk.

**Windows support?**
Functional Go core + PowerShell completions + the VS Code extension. No
Ctrl-T keybinding (terminal-specific), but `suggest --inline` works
through any binding.
<!-- draft note 3 -->
<!-- draft note 11 -->
<!-- draft note 19 -->
