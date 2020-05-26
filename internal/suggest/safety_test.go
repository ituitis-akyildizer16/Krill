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

func TestEvaluateWarns(t *testing.T) {
	sc, _ := NewSafetyCheck([]string{"rm -rf /"}, []string{"rm "})
	res, err := sc.Evaluate("rm -rf build/")
	if err != nil {
		t.Fatal(err)
	}
	if res.Warning == "" {
		t.Fatal("want a warning for rm")
	}
}

func TestEvaluateDenies(t *testing.T) {
	sc, _ := NewSafetyCheck([]string{"rm -rf /"}, []string{})
	_, err := sc.Evaluate("rm -rf /")
	if err == nil {
		t.Fatal("want deny error")
	}
}
