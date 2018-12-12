package main

import (
	"github.com/globalsign/mgo/bson"
	"net/http"
)

func GetAccountDBAsBSON(query bson.M) (bson.M, error) {
	var account bson.M
	err := IAccountCollection.Find(query).One(&account)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func FetchAccountRouteHelper(acc bson.M, w http.ResponseWriter) {

}

func FetchProfileRouteHelper(acc bson.M, w http.ResponseWriter) {

}