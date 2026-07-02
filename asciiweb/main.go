package main

import (
	"asciiweb/handler"
	"log"
	"net/http"
	"os"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.Home)
	mux.HandleFunc("/download", handler.ImageDownload)

	port := os.Getenv("PORT")
	if port == "" {
		port = "7000" // fallback for local dev
	}

	log.Println("Server is running on port " + port)

	if err := http.ListenAndServe("0.0.0.0:"+port, mux); err != nil {
		log.Fatal("Failed to start server: " + err.Error())
	}
}