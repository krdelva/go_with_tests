package collections

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d, want %d, with given slice: %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("sum of slices", func(t *testing.T) {
		s1 := []int{1, 2}
		s2 := []int{8, 9}
		s3 := []int{5, 6}

		got := SumAll(s1, s2, s3)
		want := []int{3, 17, 11}

		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	t.Run("the sum of all slice tails", func(t *testing.T) {
		s1 := []int{1, 2, 3}
		s2 := []int{4, 5, 6}
		s3 := []int{7, 8, 9}

		got := SumAllTails(s1, s2, s3)
		want := []int{5, 11, 17}

		checkSums(t, got, want)
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		s1 := []int{}
		s2 := []int{1}
		s3 := []int{2, 3}

		got := SumAllTails(s1, s2, s3)
		want := []int{0, 0, 3}

		checkSums(t, got, want)
	})
}
