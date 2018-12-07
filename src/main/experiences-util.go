package main

import (
	"encoding/json"
	"github.com/globalsign/mgo/bson"
	"github.com/julienschmidt/httprouter"
	"interfaces-api"
	"interfaces-internal"
	"net/http"
)

func GetPersonalExperiencesRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params, account interface{}) {
	w.Header().Set("Content-Type", "application/json")

	if account.Type != "User" {
		SendError(w, http.StatusBadRequest, badRequest+"  (Incorrect account type! User account type required.)", 4000)
	}
	user, ok := account.(interfaces_internal.IUser)
}

func GetExperiencesRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	var user interfaces_internal.IUser

	// Find and get user object
	err := IAccountCollection.Find(bson.M{"username": p.ByName("id"), "type": "User"}).One(&user)
	if err != nil {
		if err.Error() == "not found" {
			SendError(w, http.StatusNotFound, notFound+"  (User not found? Is this the correct account type?)", 4002)
		} else {
			SendError(w, http.StatusInternalServerError, internalServerError, 4001)
		}
		return
	}

	GetExperiences(w, user)
}

func GetExperiences(w http.ResponseWriter, user interfaces_internal.IUser) {
	var ret []interfaces_api.IExperienceAPI
	for _, e := range user.Experiences { // add experiences to api array
		ret = append(ret, interfaces_api.ConvertToIExperienceAPI(e))
	}
	response, err := json.Marshal(ret) // prepare json response
	if err != nil {
		SendError(w, http.StatusInternalServerError, internalServerError, 4001)
		return
	}

	_, err = w.Write(response) // write response
	if err != nil {
		SendError(w, http.StatusInternalServerError, internalServerError, 4001)
		return
	}
}

func CreateExperienceRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params, account interface{}) {
	w.Header().Set("Content-Type", "application/json")

}

func UpdateExperienceRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params, account interface{}) {
	w.Header().Set("Content-Type", "application/json")

}

func DeleteExperienceRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params, account interface{}) {
	w.Header().Set("Content-Type", "application/json")

}

func GetExperienceValidationsRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params, account interface{}) {
	w.Header().Set("Content-Type", "application/json")

}

func ReviewExperienceValidationsRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params, account interface{}) {
	w.Header().Set("Content-Type", "application/json")

}

func EmailApproveExperienceValidationRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params, account interface{}) {
	w.Header().Set("Content-Type", "application/json")

}