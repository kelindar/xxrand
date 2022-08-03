//go:build arm64

package xxrand

import "sync/atomic"

const havex64tsc = false

// https://stackoverflow.com/questions/27693145/rdtscp-versus-rdtsc-cpuid
func x64tsc() uint64 {
	return atomic.AddUint64(&epoch, 1)
}
