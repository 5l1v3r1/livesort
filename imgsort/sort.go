package main

import (
	"flag"
	"fmt"

	"github.com/unixpickle/essentials"
)

func Sort(args []string) {
	var sorterPath string

	fs := flag.NewFlagSet("sort", flag.ExitOnError)
	fs.StringVar(&sorterPath, "data", "sort_data", "where to read sort results")
	fs.Parse(args)

	sorter, err := ReadSorter(sorterPath)
	if err != nil {
		essentials.Die(err)
	}

	sorted := sorter.Sort()
	for _, x := range sorted {
		fmt.Println(x)
	}
}
