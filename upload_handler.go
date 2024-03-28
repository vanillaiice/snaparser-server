package snaparser_server

import (
	"archive/zip"
	"bytes"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/vanillaiice/snaparser/parser"
)

// UploadHandler handles the HTTP POST request for uploading and parsing a file.
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "error parsing form", http.StatusBadRequest)
		log.Println(err)
		return
	}

	f, h, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "error getting file", http.StatusBadRequest)
		return
	}

	contentType := h.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "application/json") {
		http.Error(w, "file not json", http.StatusBadRequest)
		return
	}

	data, err := parser.ParseAll(f)
	if err != nil {
		http.Error(w, "error parsing file", http.StatusInternalServerError)
		return
	}

	if err = f.Close(); err != nil {
		http.Error(w, "error closing file", http.StatusInternalServerError)
		return
	}

	buf := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buf)
	for k, v := range data {
		replaceSlash(&k)
		temp, err := zipWriter.Create(k + ".txt")
		if err != nil {
			http.Error(w, "error creating zip archive", http.StatusInternalServerError)
			return
		}

		if err = writeContent(temp, v); err != nil {
			http.Error(w, "error parsing content", http.StatusInternalServerError)
			return
		}
	}

	if err = zipWriter.Close(); err != nil {
		http.Error(w, "error closing zip file", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Disposition", "attachement; filename=chats.zip")
	w.Header().Add("Content-Type", "application/zip")

	http.ServeContent(w, r, "chats.zip", time.Now(), bytes.NewReader(buf.Bytes()))
}
