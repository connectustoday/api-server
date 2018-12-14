/*
 *
 *     Copyright (C) 2018 ConnectUS
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU General Public License for more details.
 *
 *     You should have received a copy of the GNU General Public License
 *     along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/schema"
	"log"
	"net/http"
	"reflect"
	"strings"
)

var (
	schemaDecoder = schema.NewDecoder()
)

func init() {
	schemaDecoder.IgnoreUnknownKeys(true)
}

// Create omit fields list for VerifyFieldsExist()

func FormOmit(omitFields []string) (ret map[string]bool) {
	ret = make(map[string]bool)
	for _, e := range omitFields {
		ret[e] = true
	}
	return
}

// Pass in reference to obj

func VerifyFieldsExist(obj interface{}, omitFields map[string]bool, fillEmpty bool) bool {
	v := reflect.ValueOf(obj).Elem()
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).IsNil() {
			if !omitFields[v.Type().Field(i).Name] {
				return false
			} else if fillEmpty {
				v.Field(i).Set(reflect.New(v.Field(i).Type().Elem()))
				//v.Field(i).Set(reflect.Zero(v.Field(i).Type())) // zero the value if needed to fill empty
			}
		}
	}
	return true
}

// Write "ok" to client

func WriteOK(w http.ResponseWriter) {
	_, err := w.Write([]byte(`{"message": "` + ok + `"}`))
	if err != nil {
		log.Println(err.Error())
	}
}

// Decode JSON or Form POST request

func DecodeRequest(r *http.Request, obj interface{}) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}

	if strings.EqualFold(r.Header.Get("Content-Type"), "application/x-www-form-urlencoded") || strings.EqualFold(r.Header.Get("Content-Type"), "application/form-data") {
		return schemaDecoder.Decode(obj, r.PostForm)
	} else {
		return json.NewDecoder(r.Body).Decode(&obj)
	}
}

// Check validity of JWT token

func CheckJWTToken(token string, secret string) (*jwt.Token, error){
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) { // Verify token authenticity
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
}

// Get claims of JWT token

func GetJWTClaims(token string, secret string) (jwt.MapClaims, error) {
	tok, err := CheckJWTToken(token, secret)
	if err != nil {
		return nil, err
	}
	if !tok.Valid {
		return nil, errors.New("invalid token")
	}
	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("not ok")
	}
	return claims, nil
}

// Helper for creating JWT token

func CreateJWTTokenHelper(secret string, expiry int64, c jwt.MapClaims) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	// create jwt token for organization verification
	c["exp"] = expiry
	token.Claims = c
	tokenString, err := token.SignedString([]byte(secret)) // sign with secret
	if err != nil {
		return "", err
	}
	return tokenString, nil
}


// Helper for checking mongodb errors
// Sends error request body so only an error check is needed

func CheckMongoQueryError(w http.ResponseWriter, err error, notFoundMsg string, errCodeNotFound int, errCodeError int) error {
	if err != nil {
		if err.Error() == "not found" {
			SendError(w, http.StatusNotFound, notFound+notFoundMsg, errCodeNotFound)
		} else {
			SendError(w, http.StatusInternalServerError, internalServerError, errCodeError)
		}
	}
	return err
}

// Quick helper to generate AuthKey fields for accounts
// AuthKey is attached to a token to verify the password has not changed

func GenAuthKey() string {
	return bson.NewObjectId().Hex()
}