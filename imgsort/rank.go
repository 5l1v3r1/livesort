package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sort"

	"github.com/unixpickle/essentials"
	"github.com/unixpickle/livesort"
)

func Rank(args []string) {
	var sorterPath string
	var iters int
	var edgeFrac float64

	fs := flag.NewFlagSet("rank", flag.ExitOnError)
	fs.StringVar(&sorterPath, "data", "sort_data", "where to read sort results")
	fs.IntVar(&iters, "iters", 100, "number of iterations")
	fs.Float64Var(&edgeFrac, "edgefrac", 1, "fraction of edges to keep per iteration")
	fs.Parse(args)

	sorter, err := ReadSorter(sorterPath)
	if err != nil {
		essentials.Die(err)
	}

	rankings := map[interface{}]float64{}

	for i := 0; i < iters; i++ {
		for j, elem := range dropEdges(sorter, edgeFrac).Sort() {
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

func dropEdges(sorter *livesort.Sorter, keepFrac float64) *livesort.Sorter {
	if keepFrac == 1 {
		return sorter
	}
	perm := rand.Perm(len(sorter.Comparisons))
	keepNum := int(keepFrac*float64(len(sorter.Comparisons)) + 0.5)
	if keepNum < 0 || keepNum > len(sorter.Comparisons) {
		essentials.Die("bad edge fraction")
	}
	res := &livesort.Sorter{Elements: sorter.Elements}
	for _, i := range perm[:keepNum] {
		res.Comparisons = append(res.Comparisons, sorter.Comparisons[i])
	}
	return res
}

type rankedElem struct {
	Elem interface{}
	Rank float64
}
