package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func count (reader io.Reader) (lines, words, bytes int) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()
		lines++;
		words += len(strings.Fields(line))
        bytes += len(line) + 1 // add 1 for the newline that Scan drops
	}
	return
}

func main() {
	fileName := os.Args[1]

	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	lines, words, bytes := count(f)
	fmt.Printf("%7d %7d %7d %s\n", lines, words, bytes, fileName)
	f.Close()
}