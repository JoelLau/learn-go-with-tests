package arraysandslices

import (
	"reflect"
	"testing"
)

func BenchmarkSumAll(b *testing.B) {
    for i := 0; i < b.N; i++ {
        SumAll([]int{1, 2, 3, 4, 5}, []int{5, 6, 7, 8, 9})
    }
}

func BenchmarkSumAll2(b *testing.B) {
    for i := 0; i < b.N; i++ {
        SumAll2([]int{1, 2, 3, 4, 5}, []int{5, 6, 7, 8, 9})
    }
}

func TestSumAll(t *testing.T) {
	t.Run("sums 1 slice", func (t *testing.T) {
		firstSlice := []int{1, 1, 1}

		have := SumAll(firstSlice)
		want := []int{3}

		assertSlices(t, have, want)
	})

	t.Run("sums 3 slices", func (t *testing.T) {
		firstSlice := []int{1, 1, 1}
		secondSlice := []int{2, 2, 2}
		thirdSlice:= []int{3, 3, 3}

		have := SumAll(firstSlice, secondSlice, thirdSlice)
		want := []int{3, 6, 9}

		assertSlices(t, have, want)
	})
}

func assertSlices(t testing.TB, have []int, want []int) {
	t.Helper()

	if !reflect.DeepEqual(have, want) {
		t.Errorf("have %q want %q", have, want)
	}
}
