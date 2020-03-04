package routes

import (
	"gowiki/constants"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

//ServeVideo route that serves a video
func ServeVideo(w http.ResponseWriter, r *http.Request) {
	slices := strings.Split(r.URL.Path, "/")
	sliceLength := len(slices)
	if sliceLength < 3 || slices[2] == "" {
		http.NotFound(w, r)
		return
	}
	fileName := constants.FileServer + slices[2] + ".mp4"

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		http.NotFound(w, r)
		return
	}

	video, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer video.Close()

	http.ServeContent(w, r, fileName, time.Now(), video)
}
