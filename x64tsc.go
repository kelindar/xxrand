//go:build !arm64

package xxrand

const havex64tsc = true

// https://stackoverflow.com/questions/27693145/rdtscp-versus-rdtsc-cpuid
func x64tsc() uint64
