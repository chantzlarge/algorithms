package algorithms

// InsertionSort ...
func InsertionSort(A []int) []int {
	for j := 1; j < len(A); j++ {
		key := A[j]
		// Insert A[j] into the sorted sequence A[1..j-1].
		i := j - 1
		for i >= 0 && A[i] > key {
			A[i+1] = A[i]
			i--
		}
		A[i+1] = key
	}
	return A
}
