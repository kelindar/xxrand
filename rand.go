// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package xxrand

import (
	"math/bits"
	"runtime"
	"sync/atomic"
)

// --------------------------------- 32-bit integers ---------------------------------

// Int32 returns a tread-safe, non-cryptographic pseudorandom int32.
func Int32() int32 {
	return int32(xxhash32(uint32(next()), 0) >> 1)
}

// Int31n returns, as an int32, a non-negative pseudo-random number in the half-open interval [0,n)
// It panics if n <= 0.
func Int31n(n int32) int32 {
	if n <= 0 {
		panic("invalid argument to Int31n")
	}

	return int32(xxhash32(uint32(next()), 0) % uint32(n))
}

// Uint32 returns a tread-safe, non-cryptographic pseudorandom uint32.
func Uint32() uint32 {
	return xxhash32(uint32(next()), 0)
}

// Uint32n returns a tread-safe, non-cryptographic pseudorandom uint32 in the range [0..maxN).
func Uint32n(maxN uint32) uint32 {
	return xxhash32(uint32(next()), 0) % uint32(maxN)
}

// --------------------------------- 64-bit integers ---------------------------------

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64
func Int63() int64 {
	return int64(xxhash64(next(), 0) >> 1)
}

// Int63n returns, as an int64, a non-negative pseudo-random number in the half-open interval [0,n). It panics if n <= 0.
func Int63n(n int64) int64 {
	if n <= 0 {
		panic("invalid argument to Int63n")
	}

	return int64(xxhash64(next(), 0) % uint64(n))
}

// Uint64 returns a tread-safe, non-cryptographic pseudorandom uint64.
func Uint64() uint64 {
	return xxhash64(next(), 0)
}

// Uint64n returns a tread-safe, non-cryptographic pseudorandom uint64 in the range [0..maxN).
func Uint64n(max uint64) uint64 {
	return xxhash64(next(), 0) % uint64(max)
}

// --------------------------------- Misc Types ---------------------------------

// Intn returns, as an int, a non-negative pseudo-random number in the half-open interval [0,n).
// It panics if n <= 0.
func Intn(n int) int {
	if n <= 0 {
		panic("invalid argument to Intn")
	}
	if n <= 1<<31-1 {
		return int(Int31n(int32(n)))
	}
	return int(Int63n(int64(n)))
}

// Bool returns, as a bool, a pseudo-random number
func Bool() bool {
	return Uint32n(2) == 0
}

// Float32 returns, as a float32, a pseudo-random number in [0.0,1.0)
func Float32() float32 {
again: // Source: math/rand. Copyright 2009 The Go Authors.
	f := float32(Float64())
	if f == 1 {
		goto again
	}
	return f
}

// Float64 returns, as a float64, a pseudo-random number in the half-open interval [0.0,1.0).
func Float64() float64 {
again: // Source: math/rand. Copyright 2009 The Go Authors.
	f := float64(Int63()) / (1 << 63)
	if f == 1 {
		goto again
	}
	return f
}

// --------------------------------- Hashing ---------------------------------

// https://stackoverflow.com/questions/27693145/rdtscp-versus-rdtsc-cpuid
func x64tsc() uint64

// Next returns the next epoch for the random, when called without an explicit
// epoch provided. This one simply reads a time stamp counter if available and
// uses it as the epoch.
var next func() uint64
var epoch uint64

// genericNext provides a cross-platform implementation, but uses atomic counter instead.
func genericNext() uint64 {
	return atomic.AddUint64(&epoch, 1)
}

func init() {
	next = genericNext
	if runtime.GOARCH == "amd64" {
		next = x64tsc
	}
}

// The unrolled xxhash that hashes the input uint32. It produces the exact same output
// as xxh3, which has a good overall distribution and a passing chi2 test.
func xxhash32(v, seed uint32) uint32 {
	return uint32(xxhash64((uint64(v) + uint64(v)<<32), uint64(seed)))
}

// The unrolled xxhash that hashes the input uint32. It produces the exact same output
// as xxh3, which has a good overall distribution and a passing chi2 test.
func xxhash64(v, seed uint64) uint64 {

	// Source: https://github.com/zeebo/xxh3
	x := v ^ (0x1cad21f72c81017c ^ 0xdb979083e96dd4de) + seed
	x ^= bits.RotateLeft64(x, 49) ^ bits.RotateLeft64(x, 24)
	x *= 0x9fb21c651e98df25
	x ^= (x >> 35) + 4
	x *= 0x9fb21c651e98df25
	x ^= (x >> 28)
	return x
}
