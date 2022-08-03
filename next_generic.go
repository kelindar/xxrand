//go:build !amd64

package xxrand

import "sync/atomic"

// Next returns the next epoch for the random, when called without an explicit
// epoch provided. This one simply reads a time stamp counter if available and
// uses it as the epoch.
var epoch uint64

// genericNext provides a cross-platform implementation, but uses atomic counter instead.
func next() uint64 {
	return atomic.AddUint64(&epoch, 1)
}
