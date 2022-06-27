// Code generated by command: go run sum_asm.go -out ../../matfuncs/sum_amd64.s -stubs ../../matfuncs/sum_amd64_stubs.go -pkg matfuncs. DO NOT EDIT.

//go:build amd64 && gc && !purego

package matfuncs

// SumAVX32 returns the sum of all values of x (32 bits, AVX required).
//go:noescape
func SumAVX32(x []float32) float32

// SumAVX64 returns the sum of all values of x (64 bits, AVX required).
//go:noescape
func SumAVX64(x []float64) float64

// SumSSE32 returns the sum of all values of x (32 bits, SSE required).
//go:noescape
func SumSSE32(x []float32) float32

// SumSSE64 returns the sum of all values of x (64 bits, SSE required).
//go:noescape
func SumSSE64(x []float64) float64
