package main

import "fmt"
import "os"

func Fib(n int, f []uint64) uint64 {
	if f[n] != 0 {
		return f[n]
	}
	f[n] = Fib(n-1, f) + Fib(n-2, f)
	return f[n]
}

func main() {
	N := 50
	var f = make([]uint64, N)
	for i := 0; i < N; i++ {
		f[i] = 0
	}
	f[0] = 1
	f[1] = 2
	Fib(N-1, f)
	var S uint64 = 0
	for i := 0; i < N; i++ {
		if f[i] >= 4000000 {
			break
		}
		if f[i]%2 == 0 {
			S += f[i]
		}
	}
	name := os.Args[0]
	fmt.Printf("Answer to problem %v: %v\n", name[len(name)-3:], S)
}
