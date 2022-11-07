package Hello

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, world."

	if got := Hello(); want != got {
		t.Errorf("Hello() = %q, want = %q", got, want)
	}
}
