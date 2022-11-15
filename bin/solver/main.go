package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/jmhobbs/buzzy/internal"
)

func main() {
	var dir *string = flag.String("dir", "../../lookup", "dictionary lookup directory")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "usage: %s <options> [pivot letter] [letter 1]...[letter 6]\n", os.Args[0])
		fmt.Fprintln(flag.CommandLine.Output(), "")
		fmt.Fprintln(flag.CommandLine.Output(), "  Letters must be lowercase, a-z.")
		fmt.Fprintln(flag.CommandLine.Output(), "")
		flag.PrintDefaults()
	}
	flag.Parse()

	args := flag.Args()

	if len(args) != 7 {
		flag.Usage()
		os.Exit(1)
	}

	pivot := args[0]
	letters := args[1:]

	fmt.Println("Pivoting on:", pivot)
	fmt.Println("With letters:", strings.Join(letters, " "))

	patterns := generateCombinations(pivot, letters)

	fmt.Println("Loading from patterns...")

	all := sortableStrings{}

	table := internal.NewDiskLookupTable(*dir)

	for _, pattern := range patterns {
		words, err := table.Lookup(pattern)
		if err != nil {
			panic(err)
		}
		all = append(all, words...)
	}

	fmt.Printf("Found %d words:\n\n", len(all))

	sort.Sort(all)

	for _, word := range all {
		fmt.Printf("  %s\n", word)
	}
}

func generateCombinations(prefix string, set []string) []string {
	combinations := []string{}
	for i, chr := range set {
		combinations = append(combinations, prefix+chr)
		combinations = append(combinations, generateCombinations(prefix+chr, set[i+1:])...)
	}
	return combinations
}

type sortableStrings []string

func (s sortableStrings) Len() int {
	return len(s)
}

func (s sortableStrings) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortableStrings) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
