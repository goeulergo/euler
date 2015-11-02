package main

import "fmt"
import "os"

func pop_fib(n int, f []uint64) uint64 {
	if f[n] != 0 {
		return f[n]
	}
	f[n] = pop_fib(n-1, f) + pop_fib(n-2, f)
	return f[n]
}

func main() {
	N := 50
	var fib_values = make([]uint64, N)
	for i := 0; i < N; i++ {
		fib_values[i] = 0
	}
	fib_values[0] = 1
	fib_values[1] = 2
	pop_fib(N-1, fib_values)

	var sum uint64 = 0
	for i := 0; i < N; i++ {
		if fib_values[i] >= 4000000 {
			break
		}
		if fib_values[i]%2 == 0 {
			sum += fib_values[i]
		}
	}

	name := os.Args[0]
	fmt.Printf("Answer to problem %v: %v\n", name[len(name)-3:], sum)
}
