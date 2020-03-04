package constants

import "os"

var (
	//FileServer directory where videos are saved
	FileServer = os.Getenv("HOME") + "/server-videos/"
)

const (
	//MaxUploadSize 32MB
	MaxUploadSize = 32 << 20
)
