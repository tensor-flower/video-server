package routes

import (
	"fmt"
	"gowiki/constants"
	"gowiki/utils"
	"io"
	"net/http"
	"os"
)

//UploadFileHandler route that receives video upload and save
func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(&w)
	if r.Method != http.MethodPost {
		return
	}
	if e := r.ParseMultipartForm(constants.MaxUploadSize); e != nil {
		http.Error(w, e.Error(), http.StatusBadRequest)
		fmt.Println("parse err", e.Error())
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()
	fmt.Fprintf(w, "%v", handler.Header)
	f, err := os.OpenFile(constants.FileServer+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
}
