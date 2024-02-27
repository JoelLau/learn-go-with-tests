package arraysandslices

import "testing"

func TestSum(t *testing.T) {
	t.Run("1 + 2 + 3 + 4 + 5 = 15", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		have := Sum(numbers)
		want := 15

		assertNumbers(t, have, want)
	})

	t.Run("1 + 2 + 3 = 6", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		have := Sum(numbers)
		want := 6

		assertNumbers(t, have, want)
	})
}

func assertNumbers(t testing.TB, have, want int)  {
	t.Helper()

	if have != want {
		t.Errorf("have %q want %q", have, want)
	}
}
