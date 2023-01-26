package pkg

import (
	"fmt"

	"golang.org/x/oauth2"
)

var endpoint = &oauth2.Endpoint{
	AuthURL:  "http://localhost:30000/realms/master/protocol/openid-connect/auth",
	TokenURL: "http://localhost:30000/realms/master/protocol/openid-connect/token",
}
var oauthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:8080/auth/callback/",
	ClientID:     "test-auth",
	ClientSecret: "h3D6fHOvPjPRhp0NN3YHGcQrFTZIBIIw",
	Endpoint:     *endpoint,
}

func Run() {
	fmt.Println(oauthConfig)
}
