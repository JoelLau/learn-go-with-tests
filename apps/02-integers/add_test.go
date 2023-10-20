package integers

import "testing"
import "fmt"

func TestAdder(t *testing.T) {
	t.Run("Simple addition", func(t *testing.T) {
		have := Add(1, 2)
		want := 3

		if have != want {
			t.Errorf("have %d want %d", have, want)
		}
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
