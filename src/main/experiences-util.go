package main

import (
	"github.com/julienschmidt/httprouter"
	"interfaces-internal"
	"net/http"
)

func GetPersonalExperiencesRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params, account interfaces_internal.IAccount) {

}

func GetExperiencesRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func CreateExperienceRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params, account interfaces_internal.IAccount) {

}

func UpdateExperienceRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params, account interfaces_internal.IAccount) {

}

func DeleteExperienceRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params, account interfaces_internal.IAccount) {

}

func GetExperienceValidationsRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params, account interfaces_internal.IAccount) {

}

func ReviewExperienceValidationsRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params, account interfaces_internal.IAccount) {

}

func EmailApproveExperienceValidationRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params, account interfaces_internal.IAccount) {

}