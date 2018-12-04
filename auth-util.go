package api_server

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/globalsign/mgo/bson"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Account middleware

func WithAccountVerify(handle httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		token := r.Header.Get("x-access-token")
		if token == "" {
			SendError(w, http.StatusUnauthorized, "No token provided.", 3000)
			return
		}

		checkToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(SECRET), nil
		})

		if err == nil && checkToken.Valid {
			if claims, ok := checkToken.Claims.(jwt.MapClaims); ok {
				IAccountCollection.Find(bson.M{"username": claims["username"]})

				return
			}
		}
		SendError(w, http.StatusInternalServerError, "Failed to authenticate token.", 3002)
	}
}
