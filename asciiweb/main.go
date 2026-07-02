package main

import (
	"asciiweb/handler"
	"log"
	"net/http"
)





func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.Home)
	mux.HandleFunc("/download", handler.ImageDownload)

	log.Println("Server is running on http://localhost:7000")

	if err := http.ListenAndServe(":7000", mux); err != nil{
		log.Fatal("Failed to start server")
	}
}