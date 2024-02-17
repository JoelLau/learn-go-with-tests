package integers

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func TestAdd(t *testing.T) {
	t.Run("2 + 2 = 4", func(t * testing.T)  {
		have := Add(2, 2)
		want := 4

		assertIntegers(t, have, want)
	})

	t.Run("2 + -3 = -1", func(t * testing.T)  {
		have := Add(2, -3)
		want := -1

		assertIntegers(t, have, want)
	})

}

func assertIntegers(t testing.TB, have, want int) {
	t.Helper()

	if have != want {
		t.Errorf("have '%d' want '%d'", have, want)
	}
}

