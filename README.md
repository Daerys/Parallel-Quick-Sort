# Parallel-Quick-Sort

# Architecture Overview

```
internal/
  usecase/
    quicksort/
      interface.go                  # интерфейс QuickSort
      bench_test.go                 # бенчмарк: seq vs par
      seq/
        quicksort.go                # последовательная реализация
        quicksort_test.go           # тесты
      parallel/
        quicksort.go                # параллельная реализация
        quicksort_test.go           # тесты
```

---

# Benchmark Results

### **Sequential (1 core)**

| Run | Time (ms) |
| --- | --------- |
| 1   | 12747.60 ms   |
| 2   | 12706.83 ms   |
| 3   | 12712.92 ms   |
| 4   | 12728.57 ms   |
| 5   | 13243.51 ms   |

**Average:** `12827.89` ms

---

### **Parallel (4 cores)**

| Run | Time (ms) |
| --- | --------- |
| 1   | 3933.02 ms   |
| 2   | 4152.61 ms   |
| 3   | 3930.01 ms   |
| 4   | 4309.62 ms   |
| 5   | 3779.27 ms   |

**Average:** `4020.91` ms

---

### Final Speedup

```
speedup = SEQ_AVG / PAR_AVG = 3.19
```
---

# Запуск

```bash
go test -run TestQuicksortBenchmark -v ./internal/usecase/quicksort
```

```bash
go test ./internal/usecase/quicksort/parallel -v
```

```bash
go test ./internal/usecase/quicksort/seq -v
```

---
