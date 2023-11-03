package arraysandslices

import (
	"reflect"
	"testing"
)

func sliceDeepEquals(t testing.TB, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	t.Run("should return [2, 5] when given [[2, 3], [1, 2, 3]]", func(t *testing.T) {
		numbers := [][]int{{1, 2}, {1, 2, 3}}
		got := SumAllTails(numbers...)
		want := []int{2, 5}

		sliceDeepEquals(t, got, want)
	})

	t.Run("should return [0, 5] when given [[], [1, 2, 3]]", func(t *testing.T) {
		numbers := [][]int{{}, {1, 2, 3}}
		got := SumAllTails(numbers...)
		want := []int{0, 5}

		sliceDeepEquals(t, got, want)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("should return [3, 12]] given [[1, 2], [3, 4, 5]]", func(t *testing.T) {
		numbers := [][]int{{1, 2}, {3, 4, 5}}
		got := SumAll(numbers...)
		want := []int{3, 12}

		sliceDeepEquals(t, got, want)
	})

	t.Run("should return [10, 14] given [[10], [2, 3, 4, 5]]", func(t *testing.T) {
		numbers := [][]int{{10}, {2, 3, 4, 5}}
		got := SumAll(numbers...)
		want := []int{10, 14}

		sliceDeepEquals(t, got, want)
	})
}

func TestSum(t *testing.T) {
	t.Run("should return 15 given [1, 2, 3, 4, 5]", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %v want %v when %v", got, want, numbers)
		}
	})

	t.Run("should return 125 given [5, 15, 25, 35, 45]", func(t *testing.T) {
		numbers := []int{5, 15, 25, 35, 45}
		got := Sum(numbers)
		want := 125

		if got != want {
			t.Errorf("got %v want %v when %v", got, want, numbers)
		}
	})

	t.Run("should return 45 given [5, 15, 25]", func(t *testing.T) {
		numbers := []int{5, 15, 25}
		got := Sum(numbers)
		want := 45

		if got != want {
			t.Errorf("got %v want %v when %v", got, want, numbers)
		}
	})
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum([]int{1, 2, 3, 4, 5})
	}
}
