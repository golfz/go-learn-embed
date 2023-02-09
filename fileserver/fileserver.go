package main

import (
	"embed"
	"net/http"
)

//go:embed files/*
var files embed.FS

func main() {
	fs := http.FileServer(http.FS(files))
	http.ListenAndServe(":8080", fs)
}
