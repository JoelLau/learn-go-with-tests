package arraysandslices

import "testing"

func TestSumAllTails(t *testing.T) {
	t.Run("{1, 2, 3}, {4, 5, 6} = {6, 15}", func(t *testing.T) {
		have := SumAllTails([]int{1,2,3}, []int{4, 5, 6})
		want := []int{5, 11}

		assertSlices(t, have, want)
	})

	t.Run("empty slices return 0", func(t *testing.T) {
		have := SumAllTails([]int{})
		want := []int{0}

		assertSlices(t, have, want)
	})
}
