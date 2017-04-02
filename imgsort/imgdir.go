package main

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

// ImgDir deals with a directory of images.
type ImgDir struct {
	Path  string
	Names []string
}

// NewImgDir reads the directory of images.
func NewImgDir(path string) (*ImgDir, error) {
	listing, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	var items []string
	for _, item := range listing {
		ext := strings.ToLower(filepath.Ext(item.Name()))
		if !item.IsDir() && ext == ".png" || ext == ".jpg" || ext == ".jpeg" {
			items = append(items, item.Name())
		}
	}
	return &ImgDir{Path: path, Names: items}, nil
}

// InterfaceNames returns i.Names as a []interface{}.
func (i *ImgDir) InterfaceNames() []interface{} {
	res := make([]interface{}, len(i.Names))
	for i, x := range i.Names {
		res[i] = x
	}
	return res
}

// Contains checks that the name exists.
func (i *ImgDir) Contains(name string) bool {
	for _, x := range i.Names {
		if name == x {
			return true
		}
	}
	return false
}
