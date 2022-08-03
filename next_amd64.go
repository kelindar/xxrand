//go:build amd64

package xxrand

const havex64tsc = true

// https://stackoverflow.com/questions/27693145/rdtscp-versus-rdtsc-cpuid
func next() uint64
