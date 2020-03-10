package constants

const (
	//MaxUploadSize 32MB
	MaxUploadSize = 32 << 20
	//BucketName s3
	BucketName = "video-server-sg"
	//Region is Singapore https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.RegionsAndAvailabilityZones.html
	Region = "ap-southeast-1"
	//CallbackURL is server endpoint after google oauth2 success
	CallbackURL = "https://fierce-lake-35299.herokuapp.com/callback"
	//OauthGoogleURLAPI is url for google oauth2 api
	OauthGoogleURLAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
	//ClientURL is url of react app
	ClientURL = "https://afternoon-eyrie-08330.herokuapp.com/"
)
