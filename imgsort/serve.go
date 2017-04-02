package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"

	"github.com/unixpickle/essentials"
	"github.com/unixpickle/livesort"
)

var AssetExpr = regexp.MustCompile(`[A-Za-z_]\.(js|css|svg)`)

func Serve(args []string) {
	var port int
	var dirPath string
	var savePath string
	var assetDir string

	fs := flag.NewFlagSet("serve", flag.ExitOnError)
	fs.IntVar(&port, "port", 8080, "HTTP server port")
	fs.StringVar(&dirPath, "dir", "", "path to image directory")
	fs.StringVar(&savePath, "data", "sort_data", "where to save sort results")
	fs.StringVar(&assetDir, "assets", "assets", "path to asset directory")
	fs.Parse(args)

	if dirPath == "" {
		essentials.Die("Required flag: -dir. See -help.")
	}

	log.Println("Reading image directory...")
	imgDir, err := NewImgDir(dirPath)
	if err != nil {
		essentials.Die(err)
	}

	sorter, err := initSorter(imgDir, savePath)
	if err != nil {
		essentials.Die(err)
	}

	http.HandleFunc("/pair", func(w http.ResponseWriter, r *http.Request) {
		servePair(w, sorter)
	})

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		lesser := r.FormValue("lesser")
		greater := r.FormValue("greater")
		if !imgDir.Contains(lesser) || !imgDir.Contains(greater) {
			serveError(w, r, "no such image")
			return
		}
		if !sorter.Add(lesser, greater) {
			serveError(w, r, "inconsistency detected")
			return
		}
		if err := SaveSorter(savePath, sorter); err != nil {
			serveError(w, r, "failed to save sorter: "+err.Error())
			return
		}
		servePair(w, sorter)
	})

	http.HandleFunc("/assets/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Path[len("/assets/"):]
		if !AssetExpr.Match(name) {
			name = "404.html"
		}
		http.ServeFile(w, r, filepath.Join(assetDir, name))
	})

	http.HandleFunc("/image", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		if !imgDir.Contains(name) {
			serve
		}
		http.ServeFile(w, r, filepath.Join(imgDir.Path, name))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// TODO: see if we need to check for "".
		if r.URL == "/" || r.URL == "" {
			http.ServeFile(w, r, filepath.Join(assetDir, "index.html"))
		} else {
			http.ServeFile(w, r, filepath.Join(assetDir, "404.html"))
		}
	})
}

func initSorter(imgDir *ImgDir, path string) (*livesort.Sorter, error) {
	sorter, err := ReadSorter(savePath)
	if os.IsNotExist(err) {
		log.Println("Creating new sorter...")
		sorter, err = CreateSorter(imgDir, savePath)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	} else {
		log.Println("Loaded sorter.")
	}
	return sorter, nil
}

func serveError(w http.ResponseWriter, r *http.Request, err string) {
	log.Printf("error at %s: %s", r.URL.Path, err)
	w.WriteHeader(http.StatusBadRequest)
	data, _ := json.Marshal(map[string]string{"error": err})
	w.Write(data)
}

func servePair(w http.ResponseWriter, s *livesort.Sorter) {
	name1, name2 := s.Request()
	data, _ := json.Marshal(map[string]interface{}{
		"pair": []interface{}{name1, name2},
	})
	w.Write(data)
}
