package main

import (
	"log"
	"net/http"
	"github/Tanmoy037/fileShare/internal/server"
)

func main() {
    http.HandleFunc("/upload", server.UploadHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}