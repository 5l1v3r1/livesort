// Command imgsort is a small web application for sorting
// hundreds of images based on some user-decided criteria.
package main

import (
	"encoding/gob"
	"fmt"
	"os"

	"github.com/unixpickle/essentials"
	"github.com/unixpickle/livesort"
)

func init() {
	gob.Register(&livesort.Sorter{})
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: imgsort <sub-command> [args | -help]")
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "Available commands:")
		fmt.Fprintln(os.Stderr, " serve    run HTTP server")
		fmt.Fprintln(os.Stderr, " sort     print sorted filename list")
		fmt.Fprintln(os.Stderr, " count    count the number of comparisons")
		fmt.Fprintln(os.Stderr, " rank     produce numerical rankings")
		fmt.Fprintln(os.Stderr, " numcomp  get comparison count per image")
		fmt.Fprintln(os.Stderr)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "serve":
		Serve(os.Args[2:])
	case "sort":
		Sort(os.Args[2:])
	case "count":
		Count(os.Args[2:])
	case "rank":
		Rank(os.Args[2:])
	case "numcomp":
		NumComp(os.Args[2:])
	default:
		essentials.Die("unknown command:", os.Args[1])
	}
}
