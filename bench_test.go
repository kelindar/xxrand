// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package xxrand

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

// BenchSink prevents the compiler from optimizing away benchmark loops.
var BenchSink uint64

func BenchmarkRand(b *testing.B) {
	b.Run("Uint32", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			s := uint32(0)
			for pb.Next() {
				s += Uint32()
			}
			atomic.AddUint64(&BenchSink, uint64(s))
		})
	})

	b.Run("Uint32n", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			s := uint32(0)
			for pb.Next() {
				s += Uint32n(1e6)
			}
			atomic.AddUint64(&BenchSink, uint64(s))
		})
	})

	b.Run("Float32", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			s := float32(0)
			for pb.Next() {
				s += Float32()
			}
			atomic.AddUint64(&BenchSink, uint64(s))
		})
	})

	b.Run("Uint64", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			s := uint64(0)
			for pb.Next() {
				s += Uint64()
			}
			atomic.AddUint64(&BenchSink, uint64(s))
		})
	})

	b.Run("Uint64n", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			s := uint64(0)
			for pb.Next() {
				s += Uint64n(1e6)
			}
			atomic.AddUint64(&BenchSink, uint64(s))
		})
	})

	b.Run("Float64", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			s := float64(0)
			for pb.Next() {
				s += Float64()
			}
			atomic.AddUint64(&BenchSink, uint64(s))
		})
	})

	assert.NotZero(b, BenchSink)
}

/*
cpu: Intel(R) Core(TM) i7-9700K CPU @ 3.60GHz
BenchmarkParallel/rand-001-8            394602625                3.023 ns/op           0 B/op          0 allocs/op
BenchmarkParallel/rand-008-8            360640208                3.130 ns/op           0 B/op          0 allocs/op
BenchmarkParallel/rand-032-8            368496836                3.492 ns/op           0 B/op          0 allocs/op
BenchmarkParallel/rand-128-8            392360864                3.041 ns/op           0 B/op          0 allocs/op
BenchmarkParallel/rand-512-8            388266206                3.043 ns/op           0 B/op          0 allocs/op
BenchmarkParallel/rand-2048-8           382660137                3.097 ns/op           0 B/op          0 allocs/op
BenchmarkParallel/math-001-8            34414028                35.25 ns/op            0 B/op          0 allocs/op
BenchmarkParallel/math-008-8            31955602                38.17 ns/op            0 B/op          0 allocs/op
BenchmarkParallel/math-032-8            28410368                44.16 ns/op            0 B/op          0 allocs/op
BenchmarkParallel/math-128-8            26773701                47.21 ns/op            0 B/op          0 allocs/op
BenchmarkParallel/math-512-8            34901576                46.94 ns/op            0 B/op          0 allocs/op
BenchmarkParallel/math-2048-8           28991878                40.81 ns/op            0 B/op          0 allocs/op
*/
func BenchmarkParallel(b *testing.B) {
	var result uint64
	cases := []int{1, 8, 32, 128, 512, 2048}

	for _, i := range cases {
		b.Run(fmt.Sprintf("rand-%03d", i), func(b *testing.B) {
			b.ReportAllocs()
			b.SetParallelism(i)
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				s := uint32(0)
				for pb.Next() {
					s += Uint32()
				}
				atomic.AddUint64(&result, uint64(s))
			})
		})
	}

	for _, i := range cases {
		b.Run(fmt.Sprintf("math-%03d", i), func(b *testing.B) {
			b.ReportAllocs()
			b.SetParallelism(i)
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				s := uint32(0)
				for pb.Next() {
					s += uint32(rand.Int31n(1e6))
				}
				atomic.AddUint64(&result, uint64(s))
			})
		})
	}

	assert.NotZero(b, result)
}

func BenchmarkMathRandInt31n(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		s := uint32(0)
		for pb.Next() {
			s += uint32(rand.Int31n(1e6))
		}
		atomic.AddUint64(&BenchSink, uint64(s))
	})
}
