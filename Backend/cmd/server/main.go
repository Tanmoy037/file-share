package main

import (
	"log"
	"net/http"
	"github/Tanmoy037/fileShare/internal/server"
	"fmt"
)

func main() {
    fmt.Printf("Server starting on port 8080\n")
    http.HandleFunc("/upload", server.UploadHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
