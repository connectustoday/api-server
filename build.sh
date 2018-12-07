#!/bin/bash
mkdir bin
cd src/main
go get -t github.com/julienschmidt/httprouter
go get -t github.com/globalsign/mgo
go get -t github.com/dgrijalva/jwt-go
go get -t golang.org/x/crypto/bcrypt
go get -t github.com/gorilla/schema
GOPATH=$GOPATH:$PWD/../../ go build -o api-server .
mv ./api-server ../../bin
cd ../../
