package integers

import "testing"

func TestAdder(t *testing.T) {
	t.Run("Simple addition", func(t *testing.T) {
		have := Add(1, 2)
		want := 3

		if have != want {
			t.Errorf("have %d want %d", have, want)
		}
	})
}
