package main

// みんなのGo p.34 file/filepath

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// http://localhost:8080/data/index.html で index.html を返す
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Path: " + r.URL.Path)

		// Use path package for http request
		if ok, err := path.Match("/data/*", r.URL.Path); err != nil || !ok {
			fmt.Println(ok)
			fmt.Println(err)
			http.NotFound(w, r)
			return
		}

		// Use filepath package for local(phisical) path
		name := filepath.Join(cwd, "data", filepath.Base(r.URL.Path))

		// もし誤ってpathを使うと、Windowsにおいて、
		// http://localhost:8080/..\main.go でこのファイルを返してしまう
		// name := filepath.Join(cwd, "data", path.Base(r.URL.Path))

		f, err := os.Open(name)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		defer f.Close()
		io.Copy(w, f)
	})
	http.ListenAndServe("127.0.0.1:8080", nil)
}
