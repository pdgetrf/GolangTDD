package array

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		actual := Sum(numbers)
		expect := 15
		if actual != expect {
			t.Errorf("expect: %d, actual %d, data: %v", expect, actual, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	expect := []int{3, 9}

	if !reflect.DeepEqual(got, expect) {
		t.Errorf("got %v expect %v", got, expect)
	}
}

func TestSumAllTails(t *testing.T) {
	t.Run("happy cases", func(t *testing.T) {
		got := SumAllTrails([]int{1, 2}, []int{0, 2, 9})
		want := []int{2, 11}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("single empty slice", func(t *testing.T) {
		got := SumAllTrails([]int{})
		want := []int{0}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("multiple empty slices", func(t *testing.T) {
		got := SumAllTrails([]int{}, []int{})
		want := []int{0, 0}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("mix non-empty and empty slices", func(t *testing.T) {
		got := SumAllTrails([]int{1, 2, 3}, []int{})
		want := []int{5, 0}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
