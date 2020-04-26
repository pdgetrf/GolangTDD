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
