package api_server

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/globalsign/mgo/bson"
	"github.com/julienschmidt/httprouter"
	"interfaces-internal"
	"net/http"
)

func GetAccountFromContext(ctx context.Context) interfaces_internal.IAccount {
	return ctx.Value("account").(interfaces_internal.IAccount) // Low possibility of nil (should've been addressed by middleware)
}

// Account middleware

func WithAccountVerify(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		token := r.Header.Get("x-access-token")
		if token == "" {
			SendError(w, http.StatusUnauthorized, "No token provided.", 3000)
			return
		}

		checkToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) { // Verify token authenticity
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(SECRET), nil
		})

		if err == nil && checkToken.Valid {
			if claims, ok := checkToken.Claims.(jwt.MapClaims); ok {

				query := IAccountCollection.Find(bson.M{"username": claims["username"]})

				count, err := query.Count() // Check if account exists
				if err != nil {
					SendError(w, http.StatusInternalServerError, internalServerError+" (Problem finding account)", 3002)
					return
				}
				if count == 0 {
					SendError(w, http.StatusNotFound, notFound+" (Account not found)", 3003)
					return
				}

				result := interfaces_internal.IAccount{} // Get account
				err = query.One(&result)
				if err != nil {
					if DEBUG {
						println(err)
					}
					SendError(w, http.StatusInternalServerError, internalServerError+" (Problem finding account)", 3002)
					return
				}
				if !result.IsEmailVerified { // Check for Email Verification
					SendError(w, http.StatusUnauthorized, unauthorized+" (Email not verified.)", 3004)
					return
				}

				r.Context().Value("account") = result // coolest thing about golang :OO
				r.Context().Value("accountType") = result.Type

				next(w, r, params) // call next middleware or main router function
				return
			}
		}
		SendError(w, http.StatusInternalServerError, "Failed to authenticate token.", 3002)
	}
}

func LoginRoute(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}
