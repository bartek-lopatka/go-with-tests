package slices

import (
	"reflect"
	"testing"
)

func TestSumAllTails(t *testing.T) {

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{5, 7})
		want := []int{2, 7}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})

	t.Run("safely sum with empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{5, 7, 8})
		want := []int{0, 15}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}
