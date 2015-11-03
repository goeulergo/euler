package main

import "fmt"
import "os"
import "strings"
import "strconv"
import "github.com/golang/example/stringutil"

func main() {

	N := 999*999
	var x int
	Outer:
	for n := N; n > 10001; n-- {
		if strings.EqualFold(stringutil.Reverse(strconv.Itoa(n)), strconv.Itoa(n)) {
			for q := 999; q >= 100; q-- {
				if n%q == 0 {
					if n/q >= 100 && n/q < 1000 {
						x = n
						break Outer
					}
				}
			}
		}
	}

	name := os.Args[0]
	fmt.Printf("Answer to problem %v: %v\n", name[len(name)-3:], x)
}
