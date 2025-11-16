package quicksort_test

import (
	"fmt"
	"quicksort/v0/internal/usecase/quicksort/parallel"
	"quicksort/v0/internal/usecase/quicksort/seq"
	"runtime"
	"testing"
	"time"
)

const benchSize = 100_000_000
const runs = 5

func generateBase() []int32 {
	base := make([]int32, benchSize)
	seed := int64(42)
	r := newRandSource(seed)
	for i := range base {
		base[i] = int32(r.Intn(2_000_000) - 1_000_000)
	}
	return base
}

type randSource struct{ x uint64 }

func newRandSource(seed int64) *randSource {
	return &randSource{x: uint64(seed)}
}
func (r *randSource) Intn(n int) int {
	r.x ^= r.x << 13
	r.x ^= r.x >> 7
	r.x ^= r.x << 17
	return int(r.x % uint64(n))
}

func isSorted(a []int32) bool {
	for i := 1; i < len(a); i++ {
		if a[i-1] > a[i] {
			return false
		}
	}
	return true
}

func TestQuicksortBenchmark(t *testing.T) {
	fmt.Println("Generating 1e8-element array...")
	base := generateBase()
	tmp := make([]int32, benchSize)

	seqSorter := seq.QuickSort{}
	parSorter := parallel.QuickSort{}

	// ---------------------------
	//     SEQUENTIAL (1 core)
	// ---------------------------
	runtime.GOMAXPROCS(1)
	fmt.Println("Running sequential benchmark (GOMAXPROCS=1)...")

	var seqSum float64
	for i := 0; i < runs; i++ {
		copy(tmp, base)
		start := time.Now()

		seqSorter.Sort(tmp)

		elapsed := time.Since(start).Seconds() * 1000
		seqSum += elapsed

		fmt.Printf("SEQ run %d: %.2f ms\n", i+1, elapsed)

		if !isSorted(tmp) {
			t.Fatalf("seq result is not sorted on run %d", i+1)
		}
	}

	seqAvg := seqSum / runs
	fmt.Printf("SEQ average: %.2f ms\n\n", seqAvg)

	// ---------------------------
	//     PARALLEL (4 cores)
	// ---------------------------
	runtime.GOMAXPROCS(4)
	fmt.Println("Running parallel benchmark (GOMAXPROCS=4)...")

	var parSum float64
	for i := 0; i < runs; i++ {
		copy(tmp, base)
		start := time.Now()

		parSorter.Sort(tmp)

		elapsed := time.Since(start).Seconds() * 1000
		parSum += elapsed

		fmt.Printf("PAR run %d: %.2f ms\n", i+1, elapsed)

		if !isSorted(tmp) {
			t.Fatalf("par result is not sorted on run %d", i+1)
		}
	}

	parAvg := parSum / runs
	fmt.Printf("PAR average: %.2f ms\n\n", parAvg)

	// ---------------------------
	//     Final comparison
	// ---------------------------
	fmt.Println("========== RESULT ==========")
	fmt.Printf("SEQ avg (1 core):  %.2f ms\n", seqAvg)
	fmt.Printf("PAR avg (4 cores): %.2f ms\n", parAvg)
	fmt.Printf("Speedup:           %.2fx\n", seqAvg/parAvg)
	fmt.Println("============================")

	if parAvg >= seqAvg {
		t.Fatalf("Parallel version is slower: seq=%.2fms par=%.2fms", seqAvg, parAvg)
	}
}
