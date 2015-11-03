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
	var sumOfSquares int = 0
	const N = 100
	for i := 1; i <= N; i++ {
		sumOfSquares += Pow(i, 2)
	}
	var squareOfSum int = 0
	for i := 1; i <= N; i++ {
		squareOfSum += i
	}
	squareOfSum = Pow(squareOfSum, 2)
	name := os.Args[0]
	fmt.Printf("Answer to problem %v: %v\n", name[len(name)-3:], squareOfSum-sumOfSquares)
}
