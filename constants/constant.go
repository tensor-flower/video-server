package constants

import "os"

var (
	//FileServer directory where videos are saved
	FileServer = os.Getenv("HOME") + "/server-videos/"
)

const (
	//MaxUploadSize 32MB
	MaxUploadSize = 32 << 20
	//BucketName s3
	BucketName = "video-server-sg"
	//Region is Singapore https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.RegionsAndAvailabilityZones.html
	Region = "ap-southeast-1"
)
