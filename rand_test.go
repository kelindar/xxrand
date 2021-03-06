// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package xxrand

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUint64n(t *testing.T) {
	const samples, bounds = 1000000, 100
	for start := time.Now(); time.Since(start) < 500*time.Millisecond; {

		// Generate a distribution
		m := make(map[int]float64)
		for i := 0; i < samples; i++ {
			n := Uint64n(bounds)
			if n >= bounds {
				t.Fatalf("n > %v: %v", n, bounds)
			}

			m[int(n)]++
		}

		// check distribution
		// TODO: better off with a Chi square test (http://www.stat.yale.edu/Courses/1997-98/101/chigf.htm)
		avg := float64(samples) / float64(bounds)
		for k, v := range m {
			if p := math.Abs(v-avg) / avg; p > 0.05 {
				t.Fatalf("skew more than 5%% for k=%v: %v%%", k, p*100)
			}
		}

		time.Sleep(10 * time.Microsecond)
	}
}

func TestBinary(t *testing.T) {
	const samples = 10000

	// Generate a distribution
	m := make(map[int]float64)
	for i := 0; i < samples; i++ {
		if n := Bool(); n == true {
			m[0]++
		} else {
			m[1]++
		}
	}

	// check distribution
	avg := float64(samples) / 2.0
	for k, v := range m {
		if p := math.Abs(v-avg) / avg; p > 0.05 {
			t.Fatalf("skew more than 5%% for k=%v: %v%%", k, p*100)
		}
	}
}

func TestNext(t *testing.T) {
	assert.Greater(t, int(next()), 0)
}

func TestIntn(t *testing.T) {
	assert.Panics(t, func() {
		Intn(-1)
	})

	assert.Less(t, Intn(10), 10)
	assert.Less(t, Intn(1e12), int(1e12))
}

func TestInt31n(t *testing.T) {
	assert.Panics(t, func() {
		Int31n(-1)
	})

	assert.Less(t, int(Int31n(10)), 10)
	assert.Less(t, int(Int31n(1e9)), int(1e9))
}

func TestInt63n(t *testing.T) {
	assert.Panics(t, func() {
		Int63n(-1)
	})

	assert.Less(t, int(Int63n(10)), 10)
	assert.Less(t, int(Int63n(1e9)), int(1e9))
}

func TestInt32(t *testing.T) {
	m := make(map[int32]struct{})
	for i := 0; i < 1e3; i++ {
		n := Int32()
		assert.GreaterOrEqual(t, n, int32(0))
		if _, ok := m[n]; ok {
			assert.Fail(t, "number %v already exists", n)
		}
		m[n] = struct{}{}
	}
}

func TestInt63(t *testing.T) {
	m := make(map[int64]struct{})
	for i := 0; i < 1e3; i++ {
		n := Int63()
		assert.GreaterOrEqual(t, n, int64(0))
		if _, ok := m[n]; ok {
			assert.Fail(t, "number %v already exists", n)
		}
		m[n] = struct{}{}
	}
}

func TestUint32(t *testing.T) {
	m := make(map[uint32]struct{})
	for i := 0; i < 1e3; i++ {
		n := Uint32()
		if _, ok := m[n]; ok {
			assert.Fail(t, "number %v already exists", n)
		}
		m[n] = struct{}{}
	}
}

func TestUint64(t *testing.T) {
	m := make(map[uint64]struct{})
	for i := 0; i < 1e3; i++ {
		n := Uint64()
		if _, ok := m[n]; ok {
			assert.Fail(t, "number %v already exists", n)
		}
		m[n] = struct{}{}
	}
}

func TestBool(t *testing.T) {
	var heads, tails int
	for i := 0; i < 100; i++ {
		if Bool() {
			heads++
		} else {
			tails++
		}
	}
	assert.NotZero(t, heads)
	assert.NotZero(t, tails)
}

func TestFloat32(t *testing.T) {
	assert.GreaterOrEqual(t, Float32(), float32(0))
	assert.Less(t, Float32(), float32(1))
}

func TestFloat64(t *testing.T) {
	assert.GreaterOrEqual(t, Float64(), 0.0)
	assert.Less(t, Float64(), 1.0)
}
