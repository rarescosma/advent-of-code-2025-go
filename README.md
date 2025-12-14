# 🎄 Advent of Code 2025 in Golang 🐿️

## Targets

```console
$ make

Usage: make [target] ...

Targets:
build      Build all days and place binaries under ./bin
run        Timed run of all days
hyper      Hyperfine benckmark of all days
scaffold   Scaffold a day, prefix it with DAYS=xx
clean      Nuke the ./bin dir from orbit
lint       Run linters and checkers
```

## Inputs

Inputs for each problem are expected in the `./inputs/xx.in` files, where `xx` is the 2-digit, zero-padded day number.

## Benchmarks

```console
$ inxi -Cxxx
CPU:
  Info: 16-core (6-mt/10-st) model: Intel Core Ultra 7 155H bits: 64
    type: MST AMCP smt: enabled arch: Meteor Lake rev: 4 cache: L1: 1.6 MiB
    L2: 18 MiB L3: 24 MiB
  Speed (MHz): avg: 3400 min/max: 400/4500:4800:3800:2500 cores: 1: 3400
    2: 3400 3: 3400 4: 3400 5: 3400 6: 3400 7: 3400 8: 3400 9: 3400 10: 3400
    11: 3400 12: 3400 13: 3400 14: 3400 15: 3400 16: 3400 17: 3400 18: 3400
    19: 3400 20: 3400 21: 3400 22: 3400 bogomips: 131788
  Flags-basic: avx avx2 ht lm nx pae sse sse2 sse3 sse4_1 sse4_2 ssse3 vmx

$ make hyper
+--------+
| Day 01 |
+--------+
Benchmark 1: ./bin/01
  Time (mean ± σ):       1.7 ms ±   0.1 ms    [User: 0.5 ms, System: 1.2 ms]
  Range (min … max):     1.3 ms …   2.8 ms    1070 runs


+--------+
| Day 02 |
+--------+
Benchmark 1: ./bin/02
  Time (mean ± σ):      14.5 ms ±   1.5 ms    [User: 105.8 ms, System: 4.9 ms]
  Range (min … max):    11.0 ms …  20.6 ms    164 runs


+--------+
| Day 03 |
+--------+
Benchmark 1: ./bin/03
  Time (mean ± σ):       2.8 ms ±   0.3 ms    [User: 1.4 ms, System: 1.5 ms]
  Range (min … max):     2.0 ms …   3.5 ms    1101 runs


+--------+
| Day 04 |
+--------+
Benchmark 1: ./bin/04
  Time (mean ± σ):      21.8 ms ±   2.3 ms    [User: 20.1 ms, System: 1.7 ms]
  Range (min … max):    18.1 ms …  32.6 ms    130 runs


+--------+
| Day 05 |
+--------+
Benchmark 1: ./bin/05
  Time (mean ± σ):       2.7 ms ±   0.2 ms    [User: 1.6 ms, System: 1.3 ms]
  Range (min … max):     2.1 ms …   3.3 ms    902 runs


+--------+
| Day 06 |
+--------+
Benchmark 1: ./bin/06
  Time (mean ± σ):       2.9 ms ±   0.3 ms    [User: 1.3 ms, System: 1.6 ms]
  Range (min … max):     2.0 ms …   5.9 ms    925 runs


+--------+
| Day 07 |
+--------+
Benchmark 1: ./bin/07
  Time (mean ± σ):       1.9 ms ±   0.2 ms    [User: 0.8 ms, System: 1.2 ms]
  Range (min … max):     1.4 ms …   3.1 ms    1195 runs


+--------+
| Day 08 |
+--------+
Benchmark 1: ./bin/08
  Time (mean ± σ):      99.6 ms ±   3.7 ms    [User: 95.7 ms, System: 4.6 ms]
  Range (min … max):    94.9 ms … 111.4 ms    30 runs


+--------+
| Day 09 |
+--------+
Benchmark 1: ./bin/09
  Time (mean ± σ):      86.7 ms ±   5.4 ms    [User: 89.2 ms, System: 6.4 ms]
  Range (min … max):    78.2 ms …  99.1 ms    34 runs


+--------+
| Day 10 |
+--------+
Benchmark 1: ./bin/10
  Time (mean ± σ):      63.3 ms ±   1.3 ms    [User: 63.3 ms, System: 5.8 ms]
  Range (min … max):    60.1 ms …  67.8 ms    49 runs


+--------+
| Day 11 |
+--------+
Benchmark 1: ./bin/11
  Time (mean ± σ):       2.9 ms ±   0.2 ms    [User: 1.7 ms, System: 1.3 ms]
  Range (min … max):     2.2 ms …   5.3 ms    800 runs


+--------+
| Day 12 |
+--------+
Benchmark 1: ./bin/12
  Time (mean ± σ):       2.1 ms ±   0.1 ms    [User: 0.9 ms, System: 1.2 ms]
  Range (min … max):     1.7 ms …   2.5 ms    1225 runs
```
