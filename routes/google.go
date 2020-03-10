package routes

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"os"
	"time"
	"video-server/constants"
	"video-server/utils"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOauthConfig = &oauth2.Config{
	RedirectURL:  constants.CallbackURL,
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

func generateStateOauthCookie(w http.ResponseWriter) string {
	//1 month
	var expiration = time.Now().Add(30 * 24 * time.Hour)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(w, &cookie)

	return state
}

//OauthGoogleLogin google oauth2 api
func OauthGoogleLogin(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(&w)
	oauthState := generateStateOauthCookie(w)
	u := googleOauthConfig.AuthCodeURL(oauthState)
	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}
