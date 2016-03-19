package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	s, sep := "", ""

	// inefficient method
	start := time.Now()
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	fmt.Println("Elapsed:", time.Since(start))
	fmt.Println(s)

	// using Join
	start = time.Now()
	s = strings.Join(os.Args, sep)
	fmt.Println("Elapsed:", time.Since(start))
	fmt.Println(s)
}
