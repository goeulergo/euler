package main

import "fmt"
import "os"

func main() {
	sum := 0
	for n := 1; n < 1000; n++ {
		if n%3 == 0 || n%5 == 0 {
			sum += n
		}
	}
	name := os.Args[0]
	fmt.Printf("Answer to problem %v: %v\n", name[len(name)-3:], sum)
}
