package main

import (
	"github.com/globalsign/mgo/bson"
	"github.com/julienschmidt/httprouter"
	"interfaces-internal"
	"net/http"
)

func GetPersonalExperiencesRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params, account interfaces_internal.IAccount) {
	w.Header().Set("Content-Type", "application/json")

}

func GetExperiencesRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

}

func GetExperiences(w http.ResponseWriter, userName string) {
	var user interfaces_internal.IUser
	err := IAccountCollection.Find(bson.M{"username": userName, "type": "User"}).One(&user)
	if err != nil {
		if err.Error() == "not found" {
			SendError(w, http.StatusNotFound, notFound+" (Account not found)", 3003)
		} else {
			SendError(w, http.StatusInternalServerError, internalServerError+" (Problem finding account)", 3002)
		}
		return
	}
}

func CreateExperienceRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params, account interfaces_internal.IAccount) {
	w.Header().Set("Content-Type", "application/json")

}

func UpdateExperienceRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params, account interfaces_internal.IAccount) {
	w.Header().Set("Content-Type", "application/json")

}

func DeleteExperienceRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params, account interfaces_internal.IAccount) {
	w.Header().Set("Content-Type", "application/json")

}

func GetExperienceValidationsRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params, account interfaces_internal.IAccount) {
	w.Header().Set("Content-Type", "application/json")

}

func ReviewExperienceValidationsRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params, account interfaces_internal.IAccount) {
	w.Header().Set("Content-Type", "application/json")

}

func EmailApproveExperienceValidationRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params, account interfaces_internal.IAccount) {
	w.Header().Set("Content-Type", "application/json")

}