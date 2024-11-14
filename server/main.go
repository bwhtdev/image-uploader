package main

import (
	"fmt"
	"net/http"

	"github.com/olahol/go-imageupload"
)

func main() {
  fs := http.FileServer(http.Dir("../images/"))
  http.Handle("/image/", http.StripPrefix("/image/", fs))

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../index.html")
	})

	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.NotFound(w, r)
			return
		}

		img, err := imageupload.Process(r, "file")

		if err != nil {
			panic(err)
		}

		thumb, err := imageupload.ThumbnailPNG(img, 300, 300)

		if err != nil {
			panic(err)
		}

    name := img.Filename

		thumb.Save(fmt.Sprintf("../images/%s", name))

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	})

	http.ListenAndServe(":5000", nil)
}
