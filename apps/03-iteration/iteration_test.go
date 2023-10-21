package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("given 'a' returns 'aaaaa'", func(t *testing.T) {
		have := Repeat("a")
		want := "aaaaa"

		if have != want {
			t.Errorf("have %q want %q", have, want)
		}
	})

	t.Run("given 'b' returns 'bbbbb'", func(t *testing.T) {
		have := Repeat("b")
		want := "bbbbb"

		if have != want {
			t.Errorf("have %q want %q", have, want)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
