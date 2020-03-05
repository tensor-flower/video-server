package routes

import (
	"fmt"
	"net/http"
	"video-server/constants"
	"video-server/utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
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
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	defer file.Close()

	// The session the S3 Uploader will use
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(constants.Region),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)
	var bucketName = constants.BucketName
	var keyName = handler.Filename
	upParams := &s3manager.UploadInput{
		Bucket: &bucketName,
		Key:    &keyName,
		Body:   file,
	}

	// Perform an upload
	result, err := uploader.Upload(upParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
