package main

import "fmt"
import "os"
import "math"

func main() {
	T := 600851475143
	N := int(math.Sqrt(float64(T))) + 1

	sieve := make([]bool, N)

	for i := 0; i < N; i++ {
		sieve[i] = true
	}
	for i := 2; i < N; i++ {
		if sieve[i] {
			for p := 2 * i; p < N; p += i {
				sieve[p] = false
			}
		}
	}
	var P int
	for i := N - 1; i >= 0; i-- {
		if sieve[i] && T%i == 0 {
			P = i
			break
		}
	}

	name := os.Args[0]
	fmt.Printf("Answer to problem %v: %v\n", name[len(name)-3:], P)
}
