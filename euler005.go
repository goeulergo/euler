package main

import "fmt"
import "os"

func Pow(p int, q int) int {
	var r int = 1
	for i := 0; i < q; i++ {
		r*=p
	}
	return r
}

func main() {
	const N = 21
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
	var factors [N][N]int
	for i := 2; i < N; i++ {
		for p := 2; p < N; p++ {
			if sieve[p] {
				r := i
				for r%p == 0 {
					factors[i][p]++;
					r = r/p
				}
			}
		}
	}

	var maxfactors [N]int
	for i := 2; i < N; i++ {
		for  n, q := range factors[i] {
			if maxfactors[n] < q {
				maxfactors[n] = q
			}
		}
	}

	res := 1
	for i, q := range maxfactors {
		if q != 0 {
			res *= Pow(i, q)
		}
	}

	name := os.Args[0]
	fmt.Printf("Answer to problem %v: %v\n", name[len(name)-3:], res)
}
