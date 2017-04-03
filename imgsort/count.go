package main

import (
	"flag"
	"fmt"

	"github.com/unixpickle/essentials"
)

func Count(args []string) {
	var sorterPath string

	fs := flag.NewFlagSet("sort", flag.ExitOnError)
	fs.StringVar(&sorterPath, "data", "sort_data", "where to read sort results")
	fs.Parse(args)

	sorter, err := ReadSorter(sorterPath)
	if err != nil {
		essentials.Die(err)
	}

	var relevant int
	for _, comp := range sorter.Comparisons {
		var present1, present2 bool
		for _, el := range sorter.Elements {
			present1 = present1 || (el == comp.Lesser)
			present2 = present2 || (el == comp.Greater)
		}
		if present1 && present2 {
			relevant++
		}
	}

	fmt.Printf("There are %d comparisons (%d valid)\n", len(sorter.Comparisons), relevant)
}
