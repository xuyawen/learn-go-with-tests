package integers

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}
		got := Sum(numbers)
		want := 10
		sumAssertCorrectMessage(t, got, want, numbers)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("collection of any array", func(t *testing.T) {
		arr1 := []int{1, 2, 3, 4}
		arr2 := []int{5, 6, 7}
		got := SumAll(arr1, arr2)
		want := 28
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got '%v' want '%v' given, '%v' '%v'", got, want, arr1, arr2)
		}
	})
}

func TestSumAllItem(t *testing.T) {
	t.Run("collection of any array item", func(t *testing.T) {
		got := SumAllItem([]int{1, 2, 3}, []int{4, 5, 6})
		want := []int{6, 15}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got '%v' want '%v' given", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t *testing.T, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got '%v' want '%v'", got, want)
		}
	}

	t.Run("collection of any array tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{4, 5, 6})
		want := []int{5, 11}
		checkSums(t, got, want)
	})

	t.Run("collection of empty array tails", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 3, 4})
		want := []int{0, 7}
		checkSums(t, got, want)
	})
}

func sumAssertCorrectMessage(t *testing.T, got, want int, numbers []int) {
	t.Helper()
	if got != want {
		t.Errorf("got '%d' want '%d' given, '%v'", got, want, numbers)
	}
}
