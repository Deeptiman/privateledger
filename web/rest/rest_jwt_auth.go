package rest

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("thisisafabrichyperledgerdemo")
var savedToken = make(map[string]string)

func (app *RestApp) processAuthentication(w http.ResponseWriter, key string) string {

	validToken, err := GenerateJWT(key)

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	client := &http.Client{}
	r, _ := http.NewRequest("GET", "http://localhost:"+PORT+"/token_auth", nil)
	r.Header.Set("Token", validToken)

	_, err = client.Do(r)

	if err != nil {
		respondJSON(w, map[string]string{"error": "Auth Error - "+err.Error()})
	}

	savedToken["token"] = validToken

	return validToken
}

func GenerateJWT(email string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = email
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Println("something went wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func (app *RestApp) hasSavedToken(endpoint func(http.ResponseWriter, *http.Request, string)) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		if savedToken["token"] != "" {
			token := savedToken["token"]
			endpoint(w, r, token)
		} else {
			respondJSON(w, map[string]string{"error": "Not Authorized"})
		}
	}
}

func (app *RestApp) isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {

			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {

				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was a error")
				}
				return mySigningKey, nil
			})

			if err != nil {
				respondJSON(w, map[string]string{"error": err.Error()})
			}

			if token.Valid {
				endpoint(w, r)
			}

		} else {
			respondJSON(w, map[string]string{"error": "Not Authorized"})
		}
	}
}
