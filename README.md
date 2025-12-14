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
  Time (mean ± σ):       1.1 ms ±   0.1 ms    [User: 0.3 ms, System: 0.9 ms]
  Range (min … max):     0.7 ms …   2.0 ms    1514 runs


+--------+
| Day 02 |
+--------+
Benchmark 1: ./bin/02
  Time (mean ± σ):      10.0 ms ±   0.8 ms    [User: 80.8 ms, System: 4.2 ms]
  Range (min … max):     8.4 ms …  13.2 ms    250 runs


+--------+
| Day 03 |
+--------+
Benchmark 1: ./bin/03
  Time (mean ± σ):       2.1 ms ±   0.3 ms    [User: 1.0 ms, System: 1.1 ms]
  Range (min … max):     1.4 ms …   4.1 ms    1424 runs


+--------+
| Day 04 |
+--------+
Benchmark 1: ./bin/04
  Time (mean ± σ):      12.2 ms ±   1.6 ms    [User: 11.1 ms, System: 1.2 ms]
  Range (min … max):    10.0 ms …  15.8 ms    239 runs


+--------+
| Day 05 |
+--------+
Benchmark 1: ./bin/05
  Time (mean ± σ):       1.9 ms ±   0.3 ms    [User: 1.0 ms, System: 0.9 ms]
  Range (min … max):     1.1 ms …   3.7 ms    1323 runs


+--------+
| Day 06 |
+--------+
Benchmark 1: ./bin/06
  Time (mean ± σ):       2.0 ms ±   0.2 ms    [User: 0.9 ms, System: 1.1 ms]
  Range (min … max):     1.3 ms …   2.7 ms    1492 runs


+--------+
| Day 07 |
+--------+
Benchmark 1: ./bin/07
  Time (mean ± σ):       1.3 ms ±   0.2 ms    [User: 0.5 ms, System: 0.9 ms]
  Range (min … max):     0.8 ms …   3.2 ms    1862 runs


+--------+
| Day 08 |
+--------+
Benchmark 1: ./bin/08
  Time (mean ± σ):      56.6 ms ±   2.4 ms    [User: 54.0 ms, System: 2.9 ms]
  Range (min … max):    53.1 ms …  63.2 ms    52 runs


+--------+
| Day 09 |
+--------+
Benchmark 1: ./bin/09
  Time (mean ± σ):      22.4 ms ±   1.7 ms    [User: 36.8 ms, System: 8.0 ms]
  Range (min … max):    18.4 ms …  28.4 ms    131 runs


+--------+
| Day 10 |
+--------+
Benchmark 1: ./bin/10
  Time (mean ± σ):      15.4 ms ±   1.1 ms    [User: 70.5 ms, System: 13.1 ms]
  Range (min … max):    12.9 ms …  21.5 ms    192 runs


+--------+
| Day 11 |
+--------+
Benchmark 1: ./bin/11
  Time (mean ± σ):       1.7 ms ±   0.2 ms    [User: 0.9 ms, System: 0.9 ms]
  Range (min … max):     1.0 ms …   2.6 ms    1372 runs


+--------+
| Day 12 |
+--------+
Benchmark 1: ./bin/12
  Time (mean ± σ):       1.5 ms ±   0.2 ms    [User: 0.6 ms, System: 0.9 ms]
  Range (min … max):     0.9 ms …   3.2 ms    1423 runs
```
