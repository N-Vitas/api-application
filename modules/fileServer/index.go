package fileServer

import "net/http"

func NewFileServer(patern string, path string) {
	fs := http.FileServer(http.Dir(path))
	http.Handle(patern, http.StripPrefix(patern, fs))
}