package pkg

import (
	"fmt"
	"log"

	"github.com/golang-jwt/jwt/v4"
)

func ParseJWT(token string) {
	// sample token string taken from the New example
	parser := &jwt.Parser{UseJSONNumber: true, SkipClaimsValidation: true}
	decodedToken, _, err := parser.ParseUnverified(token, jwt.MapClaims{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("decoded token: %v\n", decodedToken)
}
