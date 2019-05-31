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
author-mail <bob@example.com>
author-time 1710000000
author-tz +0000
	return 0
}
`
	lines, err := parseBlame(out, 0)
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 2 {
		t.Fatalf("want 2 lines, got %d", len(lines))
	}
	if lines[0].Commit[:8] != "4c8a1f3a" {
		t.Errorf("first commit = %q", lines[0].Commit)
	}
	if lines[0].Author != "Alice" {
		t.Errorf("first author = %q", lines[0].Author)
	}
	if lines[0].Content != "func retry() {" {
		t.Errorf("first content = %q", lines[0].Content)
	}
	if lines[1].Commit[:8] != "b91e0aa1" {
		t.Errorf("second commit = %q", lines[1].Commit)
	}
}

func TestParseBlameMaxLines(t *testing.T) {
	out := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa 1 1\nauthor A\n\t1\nbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb 2 2\nauthor B\n\t2\ncccccccccccccccccccccccccccccccccccccccc 3 3\nauthor C\n\t3\n"
	lines, err := parseBlame(out, 2)
