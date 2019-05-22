package git

import "testing"

func TestParseBlame(t *testing.T) {
	out := `4c8a1f3a2b2c3d4e5f6a7b8c9d0e1f2a3b4c5d6e 1 1
author Alice
author-mail <alice@example.com>
author-time 1700000000
author-tz +0000
	func retry() {
b91e0aa1f2b3c4d5e6f7a8b9c0d1e2f3a4b5c6d7e8f 2 2
author Bob
