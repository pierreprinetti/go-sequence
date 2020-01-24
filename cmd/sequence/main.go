package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pierreprinetti/go-sequence"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] == "version" {
		fmt.Fprintln(os.Stderr, "sequence github.com/pierreprinetti/go-sequence")
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "Pass a sequence of positive integers to get the corresponding space-separated list of numbers")
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "Example:")
		fmt.Fprintln(os.Stderr, "\tsequence '1-3,7,5-2' # 1 2 3 7 5 4 3 2")
		os.Exit(0)
	}
	s, err := sequence.Int(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}

	out := make([]string, len(s))
	for i, n := range s {
		out[i] = strconv.Itoa(n)
	}
	fmt.Println(strings.Join(out, " "))
}
