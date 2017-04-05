package main

import (
	"flag"
	"fmt"

	"github.com/unixpickle/essentials"
	"github.com/unixpickle/livesort"
)

func NumComp(args []string) {
	var sorterPath string

	fs := flag.NewFlagSet("absent", flag.ExitOnError)
	fs.StringVar(&sorterPath, "data", "sort_data", "where to read sort results")
	fs.Parse(args)

	sorter, err := ReadSorter(sorterPath)
	if err != nil {
		essentials.Die(err)
	}

	for _, elem := range sorter.Elements {
		fmt.Println(numComparisons(sorter, elem), elem)
	}
}

func numComparisons(s *livesort.Sorter, elem interface{}) int {
	var n int
	for _, c := range s.Comparisons {
		if c.Lesser == elem || c.Greater == elem {
			n++
		}
	}
	return n
}
