package suggest

import "testing"

func TestEvaluateAllows(t *testing.T) {
	sc, err := NewSafetyCheck([]string{"rm -rf /", "mkfs"}, []string{"rm ", "drop "})
	if err != nil {
		t.Fatal(err)
	}
