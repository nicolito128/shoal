package shoal_test

import (
	"slices"
	"testing"

	"github.com/nicolito128/shoal"
)

func TestMap(t *testing.T) {
	slice := []int{1, 2, 4, 8, 16, 32, 64}
	got := shoal.Map(slice, func(i int, value int) int {
		return value * value
	})

	want := []int{1, 4, 16, 64, 256, 1024, 4096}
	if !slices.Equal(want, got) {
		t.Errorf("map test failed, want = %v got = %v", want, got)
	}
}

func TestFilter(t *testing.T) {
	{
		slice := []string{"Anna", "Samuel", "Oscar", "Roberto", "Nicolas", "John"}
		got := shoal.Filter(slice, func(i int, value string) bool {
			return len(value) > 4
		})

		want := []string{"Samuel", "Oscar", "Roberto", "Nicolas"}
		if !slices.Equal(want, got) {
			t.Errorf("filter test failed, want = %v got = %v", want, got)
		}
	}
	{
		slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		got := shoal.Filter(slice, func(i, value int) bool {
			return value%2 != 0
		})

		want := []int{1, 3, 5, 7, 9}
		if !slices.Equal(want, got) {
			t.Errorf("filter test failed, want = %v got = %v", want, got)
		}
	}
}

func TestUnique(t *testing.T) {
	slice := []string{"A", "A", "B", "B", "C", "D", "D"}
	got := shoal.Unique(slice)

	want := []string{"A", "B", "C", "D"}
	if !slices.Equal(want, got) {
		t.Errorf("unique test failed, want = %v got = %v", want, got)
	}
}

func TestForEach(t *testing.T) {
	slice := make([]byte, 16)
	got := 0
	shoal.ForEach(slice, func(value byte) {
		got += 1
	})

	want := 16
	if got != want {
		t.Errorf("forEach test failed, want = %v got = %v", want, got)
	}
}

func TestCount(t *testing.T) {
	slice := []int{1, 0, 2, 0, -1, 3, 1, 5, 6, 7, 8, 0}
	got := shoal.Count(slice, 0)

	want := 3
	if got != want {
		t.Errorf("count test failed, want = %v got = %v", want, got)
	}
}

func TestCountBy(t *testing.T) {
	slice := []int{4, 8, 1, 2, 2, 0, 32, 57, 34, 17, 16}
	got := shoal.CountBy(slice, func(i, value int) bool {
		return value%4 == 0 && value != 0
	})

	want := 4
	if got != want {
		t.Errorf("countBy test failed, want = %v got = %v", want, got)
	}
}
