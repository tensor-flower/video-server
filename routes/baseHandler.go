package routes

import (
	"encoding/json"
	"fmt"
	"gowiki/constants"
	"gowiki/utils"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//BaseHandler route lists all existing videos
func BaseHandler(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	videos := readDir()
	urlsJSON, _ := json.Marshal(videos)
	fmt.Fprintf(w, string(urlsJSON))
}

func readDir() []string {
	files, err := ioutil.ReadDir(constants.FileServer)
	if err != nil {
		log.Fatal(err)
	}

	var videos []string
	for _, f := range files {
		videos = append(videos, strings.Split(f.Name(), ".")[0])
	}
	return videos
}
