package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("iteration test", func(t *testing.T) {
		repeated := Repeat("a")
		expected := "aaaaa"

		if expected != repeated {
			t.Errorf("expected %q, but got %q", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a")
	}
}
