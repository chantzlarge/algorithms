package algorithms

import "testing"

func TestInsertionSort(t *testing.T) {
	A := []int{5, 2, 4, 6, 1, 3}
	got := InsertionSort(A)
	t.Logf("got:\t%v", got)
}
