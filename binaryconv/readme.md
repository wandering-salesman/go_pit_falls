# Binary conversion

run benchmark as

```bash
go test -bench . -cpu 1 -benchmem -count=10 > raw_16kb.txt
```

beautify output with [benchstat](https://pkg.go.dev/golang.org/x/perf/cmd/benchstat):

```
# install if needed
go install golang.org/x/perf/cmd/benchstat@latest

# run to get output
benchstat raw_16kb.txt
```

---

#### (output formatted using chatgpt 👉🏻👈🏻)

### 🧠 **System Stats**

```text
OS:     Fedora Linux 42 (Workstation Edition) x86_64
Kernel: Linux 6.15.9-201.fc42.x86_64
Shell:  bash 5.2.37
CPU:    12th Gen Intel(R) Core(TM) i5-12500H (16 cores) @ 4.50 GHz
Memory: 5.84 GiB / 7.44 GiB (79%)
Swap:   5.44 GiB / 16.00 GiB (34%)
```

---

### ⚙️ **Benchmark Environment**

```text
GOOS:       linux
GOARCH:     amd64
Package:    binaryconv
CPU:        12th Gen Intel(R) Core(TM) i5-12500H
Used CPU:   1
```

---

### 📊 **Performance Comparison (sec/op)**

| Function   | 4KB Payload   | 16KB Payload   | Change (%) | p-value |
| ---------- | ------------- | -------------- | ---------- | ------- |
| BinaryRead | 390.2 ns ± 2% | 1392.0 ns ± 2% | +256.74%   | p=0.000 |
| UnsafeCast | 50.13 ns ± 1% | 174.80 ns ± 2% | +248.69%   | p=0.000 |
| Manual     | 132.3 ns ± 3% | 530.6 ns ± 4%  | +300.87%   | p=0.000 |

---

### 📦 **Memory Allocation (B/op)**

| Function   | 4KB Payload    | 16KB Payload    | Change (%)     | p-value |
| ---------- | -------------- | --------------- | -------------- | ------- |
| BinaryRead | 4.750 KiB ± 0% | 18.000 KiB ± 0% | +278.95%       | p=0.000 |
| UnsafeCast | 0.000 B ± 0%   | 0.000 B ± 0%    | \~ (no change) | p=1.000 |
| Manual     | 0.000 B ± 0%   | 0.000 B ± 0%    | \~ (no change) | p=1.000 |

---

### 🔁 **Allocations (allocs/op)**

| Function   | 4KB Payload | 16KB Payload | Change         | p-value |
| ---------- | ----------- | ------------ | -------------- | ------- |
| BinaryRead | 1.000 ± 0%  | 1.000 ± 0%   | \~ (no change) | p=1.000 |
| UnsafeCast | 0.000 ± 0%  | 0.000 ± 0%   | \~ (no change) | p=1.000 |
| Manual     | 0.000 ± 0%  | 0.000 ± 0%   | \~ (no change) | p=1.000 |

---

### ✅ **Summary**

- **`UnsafeCast` is fastest and zero-alloc** – suitable for high-performance, trusted data.
- **`Manual` parsing is \~2.5x slower than UnsafeCast** but still avoids allocations.
- **`BinaryRead` is the slowest and allocates memory** – fine for flexibility but not for hot paths.
- **All methods scale linearly with payload size**, but BinaryRead's overhead grows most.
