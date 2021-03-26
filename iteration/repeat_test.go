package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaa"

	if repeated != expected {
		t.Errorf("expected %s but got %s", expected, repeated)
	}
}
