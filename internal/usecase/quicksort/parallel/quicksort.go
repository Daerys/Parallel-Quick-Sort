package parallel

import (
	"quicksort/v0/internal/usecase/quicksort"
	"sync"
)

const threshold = 1 << 19

// QuickSort — параллельная реализация quicksort.
type QuickSort struct{}

var _ quicksort.QuickSort = (*QuickSort)(nil)

func (QuickSort) Sort(data []int32) {
	if len(data) < 2 {
		return
	}

	wg := new(sync.WaitGroup)

	wg.Add(1)
	go qSortPar(data, wg)

	wg.Wait()
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

func qSortPar(data []int32, wg *sync.WaitGroup) {
	if len(data) < 2 {
		wg.Done()
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

	leftLen := left - 1
	rightLen := len(data) - left

	if leftLen > threshold {
		wg.Add(1)
		go qSortPar(data[:left-1], wg)
	}
	if rightLen > threshold {
		wg.Add(1)
		go qSortPar(data[left:], wg)
	}

	if leftLen > 0 && leftLen <= threshold {
		qSort(data[:left-1])
	}
	if rightLen > 0 && rightLen <= threshold {
		qSort(data[left:])
	}

	wg.Done()
}
