
<p align="center">
<img width="330" height="110" src=".github/logo.png" border="0" alt="kelindar/xxrand">
<br>
<img src="https://img.shields.io/github/go-mod/go-version/kelindar/xxrand" alt="Go Version">
<a href="https://pkg.go.dev/github.com/kelindar/xxrand"><img src="https://pkg.go.dev/badge/github.com/kelindar/xxrand" alt="PkgGoDev"></a>
<a href="https://goreportcard.com/report/github.com/kelindar/xxrand"><img src="https://goreportcard.com/badge/github.com/kelindar/xxrand" alt="Go Report Card"></a>
<a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg" alt="License"></a>
<a href="https://coveralls.io/github/kelindar/xxrand"><img src="https://coveralls.io/repos/github/kelindar/xxrand/badge.svg" alt="Coverage"></a>
</p>

# XXH3-Based Pseudorandom Number Generator 

This package contains an experimental implementation of a [noise based pseudorandom number generator](https://www.youtube.com/watch?v=LWFzPP8ZbdU) that scales with multiple CPUs and from my benchmarks around 10x faster than using `math.Rand`. It uses [xx3](https://github.com/zeebo/xxh3) algorithm to hash a counter, the default counter being the time stamp counter (using `RDTSC` instruction). That being said, most of use-cases probably are better of with using `math.Rand` since it should provide better randomness characteristics.

## Features

 * Optimized to **scale on multiple CPUs** when called without the state (e.g. `Int31n()`).
 * Supports most of **math/rand** functions.
 * Roughly **10x** faster than `math/rand` equivalent.

## What is this for?

 * You can use this in the benchmarks where where `math/rand` tends to create a lock contention, especially when you are benchmarking multiple goroutines in parallel and generating some random data as you go.
 * You can use this in building games with Go, where performance is critical and potentially you are generating a lot of random numbers.
 * You can use this for use-cases that do not require extremely random numbers, such as load-balancing for example.

# Benchmark Results


```
cpu: Intel(R) Core(TM) i7-9700K CPU @ 3.60GHz
BenchmarkParallel/rand-001-8            394602625                3.023 ns/op           0 B/op          0 allocs/op
BenchmarkParallel/rand-008-8            360640208                3.130 ns/op           0 B/op          0 allocs/op
BenchmarkParallel/rand-032-8            368496836                3.492 ns/op           0 B/op          0 allocs/op
BenchmarkParallel/rand-128-8            392360864                3.041 ns/op           0 B/op          0 allocs/op
BenchmarkParallel/rand-512-8            388266206                3.043 ns/op           0 B/op          0 allocs/op
BenchmarkParallel/rand-2048-8           382660137                3.097 ns/op           0 B/op          0 allocs/op
```