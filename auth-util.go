package api_server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/globalsign/mgo/bson"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	"interfaces-internal"
	"net/http"
	"time"
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
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
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
		SendError(w, http.StatusBadRequest, badRequest+" (Bad request.)", 3103)
		return
	}

	var account interfaces_internal.IAccount
	err = IAccountCollection.Find(bson.M{"username": req.username}).One(&account) // find user in database

	if err != nil {
		if err.Error() == "not found" {
			SendError(w, http.StatusBadRequest, "Invalid login.", 3101)
		} else {
			SendError(w, http.StatusInternalServerError, internalServerError, 3100)
		}
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(req.password))
	if err != nil { // check if password is valid
		SendError(w, http.StatusBadRequest, "Invalid login.", 3101)
		return
	}

	if !account.IsEmailVerified { // check if email is verified
		SendError(w, http.StatusBadRequest, badRequest + " (Email not verified.)", 3102)
		return
	}

	// generate jwt for client
	token := jwt.New(jwt.SigningMethodHS256)

	claims := make(jwt.MapClaims)
	claims["username"] = account.UserName
	claims["exp"] = time.Now().Unix() + TOKEN_EXPIRY
	token.Claims = claims
	tokenString, err := token.SignedString(SECRET) // sign with secret
	if err != nil {
		SendError(w, http.StatusInternalServerError, internalServerError, 3100)
		return
	}

	_, err = w.Write([]byte(`{"token": ` + tokenString + `}`)) // return token to client

	if err != nil {
		println(err)
		w.WriteHeader(500)
	}
}
