package rand

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUint32(t *testing.T) {
	m := make(map[uint32]struct{})
	for i := 0; i < 1e3; i++ {
		n := Uint32()
		if _, ok := m[n]; ok {
			t.Fatalf("number %v already exists", n)
		}
		m[n] = struct{}{}
	}
}

func TestUint32n(t *testing.T) {
	const samples, bounds = 1000000, 100
	for start := time.Now(); time.Since(start) < 1*time.Second; {

		// Generate a distribution
		m := make(map[int]float64)
		for i := 0; i < samples; i++ {
			n := Uint32n(bounds)
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

		time.Sleep(1 * time.Millisecond)
	}
}

func TestBinary(t *testing.T) {
	const samples = 1000

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
