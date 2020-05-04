package suggest

import "testing"

func TestEvaluateAllows(t *testing.T) {
	sc, err := NewSafetyCheck([]string{"rm -rf /", "mkfs"}, []string{"rm ", "drop "})
	if err != nil {
		t.Fatal(err)
	}
	res, err := sc.Evaluate("git add -A && git commit -m wip")
	if err != nil {
		t.Fatalf("unexpected deny: %v", err)
	}
	if res.Warning != "" {
		t.Fatalf("unexpected warning: %q", res.Warning)
	}
}
