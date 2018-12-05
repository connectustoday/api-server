package api_server

import (
	"context"
	"encoding/json"
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

				result := interfaces_internal.IAccount{} // Get account
				err = IAccountCollection.Find(bson.M{"username": claims["username"]}).One(&result)
				if err != nil {
					if err.Error() == "not found" { // Check if account exists
						SendError(w, http.StatusNotFound, notFound+" (Account not found)", 3003)
					} else {
						if DEBUG {
							println(err)
						}
						SendError(w, http.StatusInternalServerError, internalServerError+" (Problem finding account)", 3002)
					}
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

// Login route
// POST /v1/auth/login
// https://connectustoday.github.io/api-server/api-reference#login

func LoginRoute(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	type requestForm struct {
		username string
		password string
	}

	decoder := json.NewDecoder(r.Body)
	var req requestForm
	err := decoder.Decode(&req)
	if err != nil { // Check decoding error
		SendError(w, http.StatusInternalServerError, internalServerError+" (There was a problem reading the request.)", 3100)
		return
	}
	if !VerifyFieldsExist(req, FormOmit([]string{})) { // Check request for correct fields
		SendError(w, http.StatusBadRequest, badRequest+" (Bad request.)", 3100)
		return
	}

	var account interfaces_internal.IAccount
	err = IAccountCollection.Find(bson.M{"username": req.username}).One(&account)

	if err != nil {

	}
}
