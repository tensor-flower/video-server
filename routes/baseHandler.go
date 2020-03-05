package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"video-server/constants"
	"video-server/utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

//BaseHandler route lists all existing videos
func BaseHandler(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(&w)
	w.Header().Set("Content-Type", "application/json")

	// https://docs.aws.amazon.com/sdk-for-go/api/service/s3/#S3.ListObjectsV2
	svc := s3.New(session.New(&aws.Config{
		Region: aws.String(constants.Region),
	}))
	input := &s3.ListObjectsV2Input{
		Bucket:  aws.String(constants.BucketName),
		MaxKeys: aws.Int64(10),
	}
	result, err := svc.ListObjectsV2(input)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchBucket:
				fmt.Println(s3.ErrCodeNoSuchBucket, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	var videos []string
	for _, key := range result.Contents {
		videos = append(videos, strings.Split(*key.Key, ".")[0])
	}

	urlsJSON, _ := json.Marshal(videos)
	fmt.Fprintf(w, string(urlsJSON))
}
