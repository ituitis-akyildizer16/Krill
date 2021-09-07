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
