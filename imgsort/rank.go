package main

import (
	"flag"
	"fmt"
	"sort"

	"github.com/unixpickle/essentials"
)

func Rank(args []string) {
	var sorterPath string
	var iters int

	fs := flag.NewFlagSet("rank", flag.ExitOnError)
	fs.StringVar(&sorterPath, "data", "sort_data", "where to read sort results")
	fs.IntVar(&iters, "iters", 100, "number of iterations")
	fs.Parse(args)

	sorter, err := ReadSorter(sorterPath)
	if err != nil {
		essentials.Die(err)
	}

	rankings := map[interface{}]float64{}

	for i := 0; i < iters; i++ {
		for j, elem := range sorter.Sort() {
			rankings[elem] += float64(j) / float64(iters*(len(sorter.Elements)-1))
		}
	}

	var list []rankedElem
	for elem, r := range rankings {
		list = append(list, rankedElem{Elem: elem, Rank: r})
	}
	sort.Slice(list, func(i int, j int) bool {
		return list[i].Rank < list[j].Rank
	})

	for _, item := range list {
		fmt.Println(item.Rank, item.Elem)
	}
}

type rankedElem struct {
	Elem interface{}
	Rank float64
}
