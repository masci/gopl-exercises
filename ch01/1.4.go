package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	from := make(map[string][]string)
	files := os.Args[1:]

	if len(files) == 0 {
		countlines(os.Stdin, counts, nil)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			countlines(f, counts, from)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d - %s - %v\n", n, line, from[line])
		}
	}
}

func countlines(f *os.File, counts map[string]int, from map[string][]string) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		counts[input.Text()]++
		do_insert := true
		for _, fname := range from[input.Text()] {
			if fname == f.Name() {
				do_insert = false
				break
			}
		}
		if do_insert {
			from[input.Text()] = append(from[input.Text()], f.Name())
		}
	}
}
