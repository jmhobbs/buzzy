package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/jmhobbs/buzzy/internal"
)

func main() {
	var (
		src *string = flag.String("src", "words_alpha.txt", "input file, one word per line, lowercase")
		dir *string = flag.String("dir", "../../lookup", "output directory")
	)
	flag.Parse()

	err := os.MkdirAll(*dir, 0744)
	if err != nil {
		panic(err)
	}

	f, err := os.Open(*src)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// count lines
	var lines int = 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = lines + 1
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	_, err = f.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	table := internal.NewDiskLookupTable(*dir)

	var word string
	var i int = 0
	var step int = lines / 20
	if step == 0 {
		step = 1
	}

	fmt.Printf("Processing %d words...\n", lines)
	scanner = bufio.NewScanner(f)
	for scanner.Scan() {
		word = scanner.Text()
		i = i + 1

		if i%step == 0 {
			fmt.Print(".")
		}

		// too short, game takes 4 letters or more
		if len(word) < 4 {
			continue
		}

		err = table.Append(word)
		if err != nil {
			panic(err)
		}
	}
	fmt.Print(". Done!\n")
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
