package main

import (
	"encoding/gob"
	"os"

	"github.com/unixpickle/essentials"
	"github.com/unixpickle/livesort"
)

// ReadSorter reads a sorter from a file.
func ReadSorter(path string) (*livesort.Sorter, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	dec := gob.NewDecoder(f)
	var res *livesort.Sorter
	if err := dec.Decode(&res); err != nil {
		return nil, essentials.AddCtx("decode "+path, err)
	}
	return res, nil
}

// CreateSorter creates a sorter file.
func CreateSorter(imgDir *ImgDir, path string) (*livesort.Sorter, error) {
	sorter := &livesort.Sorter{Elements: imgDir.InterfaceNames()}
	return sorter, SaveSorter(path, sorter)
}

// SaveSorter saves a sorter to a file.
func SaveSorter(path string, s *livesort.Sorter) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	enc := gob.NewEncoder(f)
	if err := enc.Encode(s); err != nil {
		return essentials.AddCtx("encode "+path, err)
	}
	return nil
}
