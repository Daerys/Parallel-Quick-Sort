package seq

import "quicksort/v0/internal/usecase/quicksort"

// QuickSort — последовательная реализация quicksort.
type QuickSort struct{}

var _ quicksort.QuickSort = (*QuickSort)(nil)

func (QuickSort) Sort(data []int32) {
	qSort(data)
}

func qSort(data []int32) {
	if len(data) < 2 {
		return
	}

	pivot := data[0]
	left, right := 1, len(data)-1

	for right >= left {
		if data[left] <= pivot {
			left++
		} else {
			data[right], data[left] = data[left], data[right]
			right--
		}
	}

	data[left-1], data[0] = data[0], data[left-1]

	qSort(data[:left-1])
	qSort(data[left:])
}
