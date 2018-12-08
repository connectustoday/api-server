package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/globalsign/mgo/bson"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	"interfaces-internal"
	"net/http"
	"time"
)

// Account middleware

func WithAccountVerify(next accountPassRoute) httprouter.Handle {
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

				var result interface{} // Get account
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
				acc, _ := result.(interfaces_internal.IAccount)
				if !acc.IsEmailVerified { // Check for Email Verification
					SendError(w, http.StatusUnauthorized, unauthorized+" (Email not verified.)", 3004)
					return
				}

				next(w, r, params, result) // call next middleware or main router function
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
	w.Header().Set("Content-Type", "application/json")

	type requestForm struct {
		Username *string `json:"username" schema:"username"`
		Password *string `json:"password" schema:"password"`
	}

	var req requestForm
	err := DecodeRequest(r, &req)
	if err != nil { // Check decoding error
		SendError(w, http.StatusInternalServerError, internalServerError+" (There was a problem reading the request.)", 3100)
		return
	}
	if !VerifyFieldsExist(&req, FormOmit([]string{}), false) { // Check request for correct fields
		SendError(w, http.StatusBadRequest, badRequest+" (Bad request.)", 3103)
		return
	}

	var account interfaces_internal.IAccount
	err = IAccountCollection.Find(bson.M{"username": *req.Username}).One(&account) // find user in database
	if err != nil {
		if err.Error() == "not found" {
			SendError(w, http.StatusBadRequest, "Invalid login.", 3101)
		} else {
			SendError(w, http.StatusInternalServerError, internalServerError, 3100)
		}
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(*req.Password)) // 1 second delay to prevent brute force
	if err != nil {                                                                      // check if password is valid
		SendError(w, http.StatusBadRequest, "Invalid login.", 3101)
		return
	}
	if !account.IsEmailVerified { // check if email is verified
		SendError(w, http.StatusBadRequest, badRequest+" (Email not verified.)", 3102)
		return
	}

	// generate jwt for client
	token := jwt.New(jwt.SigningMethodHS256)

	claims := make(jwt.MapClaims)
	claims["username"] = account.UserName
	claims["exp"] = time.Now().Add(time.Second * time.Duration(TOKEN_EXPIRY)).Unix()
	token.Claims = claims
	tokenString, err := token.SignedString([]byte(SECRET)) // sign with secret
	if err != nil {
		SendError(w, http.StatusInternalServerError, internalServerError, 3100)
		return
	}

	_, err = w.Write([]byte(`{"token": ` + tokenString + `}`)) // return token to client

	if err != nil {
		SendError(w, http.StatusInternalServerError, internalServerError, 3100)
		return
	}
}
