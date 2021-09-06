# FAQ

**Why a local model instead of a cloud API?**
Privacy, latency, and cost. A 7b coding model answers most repo questions
well enough, and your diff never leaves the machine.

**Is it really fast enough?**
Context collection is ~40ms; first token ~300ms on M1 with a quantized 7b.
