//go:build !amd64

package xxrand

const havex64tsc = false

// https://stackoverflow.com/questions/27693145/rdtscp-versus-rdtsc-cpuid
func x64tsc() uint64 {
	panic("not implemented")
}
