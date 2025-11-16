package parallel

import (
	"math/rand"
	"testing"
	"time"
)

func isSorted(a []int32) bool {
	for i := 1; i < len(a); i++ {
		if a[i-1] > a[i] {
			return false
		}
	}
	return true
}

func TestQuickSort_Sort_SmallCases(t *testing.T) {
	cases := [][]int32{
		{},
		{1},
		{1, 1, 1, 1},
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
		{3, 1, 2, 3, 2, 1},
	}

	qs := QuickSort{}

	for i, c := range cases {
		data := append([]int32(nil), c...)
		qs.Sort(data)

		if !isSorted(data) {
			t.Fatalf("case %d: array is not sorted: %v", i, data)
		}
	}
}

func TestQuickSort_Sort_Random(t *testing.T) {
	qs := QuickSort{}
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	for testNum := 0; testNum < 10; testNum++ {
		n := 1 + rnd.Intn(10_000)
		data := make([]int32, n)
		for i := range data {
			data[i] = int32(rnd.Intn(2_000_000) - 1_000_000)
		}

		qs.Sort(data)

		if !isSorted(data) {
			t.Fatalf("random test %d: array is not sorted (n=%d)", testNum, n)
		}
	}
}

func Test_qSort_SortsCorrectly(t *testing.T) {
	cases := [][]int32{
		{},
		{1},
		{2, 1},
		{1, 1, 1},
		{5, 4, 3, 2, 1},
		{3, 1, 4, 1, 5, 9, 2},
	}

	for i, c := range cases {
		data := append([]int32(nil), c...)
		qSort(data)

		if !isSorted(data) {
			t.Fatalf("qSort case %d: array is not sorted: %v", i, data)
		}
	}
}
