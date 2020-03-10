package main

import (
	"net/http"
	"os"
	"video-server/routes"
)

func main() {
	http.HandleFunc("/login", routes.OauthGoogleLogin)
	http.HandleFunc("/callback", routes.OauthGoogleCallback)
	http.HandleFunc("/upload", routes.UploadFileHandler)
	http.HandleFunc("/", routes.BaseHandler)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
