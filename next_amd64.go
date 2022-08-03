//go:build amd64

package xxrand

// https://stackoverflow.com/questions/27693145/rdtscp-versus-rdtsc-cpuid
func next() uint64
