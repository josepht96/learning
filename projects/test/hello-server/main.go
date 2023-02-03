package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Response struct {
	Status     string  `json:"status"`
	StatusCode int     `json:"statusCode"`
	Body       Message `json:"body"`
}

type Message struct {
	Message string `json:"message"`
	Time    string `json:"time"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Parse takes the token string and a function for looking up the key. The latter is especially
		// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
		// head of the token to identify which key to use, but the parsed token (head and claims) is provided
		// to the callback, providing flexibility.

		tokenString := r.Header.Get("Authorization")
		splitToken := strings.Split(tokenString, "Bearer ")
		tokenString = splitToken[1]

		// Parse takes the token string and a function for looking up the key. The latter is especially
		// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
		// head of the token to identify which key to use, but the parsed token (head and claims) is provided
		// to the callback, providing flexibility.
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			parts := strings.Split(tokenString, ".")
			method := jwt.GetSigningMethod("RS256")
			err := method.Verify(strings.Join(parts[0:2], "."), parts[2], key)
			if data.valid && err != nil {
				t.Errorf("[%v] Error while verifying key: %v", data.name, err)
			}
			if !data.valid && err == nil {
				t.Errorf("[%v] Invalid key passed validation", data.name)
			}
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(token)
		// if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// 	fmt.Println(claims["sub"])
		// } else {
		// 	fmt.Println(err)
		// }
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		dt := time.Now()
		data := Response{
			Status:     "OK",
			StatusCode: 200,
			Body: Message{
				Message: "hello from test-server",
				Time:    dt.String(),
			},
		}
		fmt.Println(data)
		json.NewEncoder(w).Encode(data)
	})
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		log.Printf("server is listening at %s", "http://localhost:8082")
		http.ListenAndServe(":8082", nil)
	}()
	wg.Wait()
}
