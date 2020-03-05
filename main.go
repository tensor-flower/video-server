package main

import (
	"net/http"
	"video-server/routes"
)

func main() {
	http.HandleFunc("/videos/", routes.ServeVideo)
	http.HandleFunc("/upload", routes.UploadFileHandler)
	http.HandleFunc("/", routes.BaseHandler)
	http.ListenAndServe(":8080", nil)
}
