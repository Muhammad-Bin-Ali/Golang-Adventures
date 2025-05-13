package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	longListing := flag.Bool("l", false, "long listing format")
	showAll := flag.Bool("a", false, "show all (hidden files)")
	flag.Parse()

	path := "."

	if flag.NArg() > 0 {
		path = flag.Arg(0)
	} else {
		path = os.Args[1]
	}

	entries, err := os.ReadDir(path)
	
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, entry := range entries {
		if !*showAll && strings.HasPrefix(entry.Name(), ".") {
			continue
		} else if *longListing {
			info, _ := entry.Info()
        	fmt.Printf("%10d %s %s\n", info.Size(), info.ModTime().Format(time.RFC822), entry.Name())
		} else {
			fmt.Println(entry.Name())
		}
	}
}