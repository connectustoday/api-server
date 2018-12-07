package main

import (
	"encoding/json"
	"github.com/globalsign/mgo/bson"
	"github.com/julienschmidt/httprouter"
	"interfaces-api"
	"interfaces-internal"
	"net/http"
)

// Get experiences of personal user (username found from token)
// GET /v1/experiences
// https://connectustoday.github.io/api-server/api-reference#experiences

func GetPersonalExperiencesRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params, account interface{}) {
	w.Header().Set("Content-Type", "application/json")

	user, ok := account.(interfaces_internal.IUser)
	if !ok || user.Type != "User" { // check if the obtained account is of user type
		SendError(w, http.StatusBadRequest, badRequest+"  (Incorrect account type! User account type required.)", 4000)
	}

	GetExperiences(w, user)
}

// Get all experiences of any user
// GET /v1/accounts/:id/experiences
// https://connectustoday.github.io/api-server/api-reference#accounts

func GetExperiencesRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	var user interfaces_internal.IUser

	// Find and get user object
	err := IAccountCollection.Find(bson.M{"username": p.ByName("id"), "type": "User"}).One(&user)
	if err != nil {
		if err.Error() == "not found" {
			SendError(w, http.StatusNotFound, notFound+"  (User not found! Is this the correct account type?)", 4002)
		} else {
			SendError(w, http.StatusInternalServerError, internalServerError, 4001)
		}
		return
	}

	GetExperiences(w, user)
}

// Get experience helper

func GetExperiences(w http.ResponseWriter, user interfaces_internal.IUser) {
	var ret []interfaces_api.IExperienceAPI
	for _, e := range user.Experiences { // add experiences to api array
		ret = append(ret, interfaces_api.ConvertToIExperienceAPI(e))
	}

	if ret == nil { // init empty slice if nothing was added
		ret = make([]interfaces_api.IExperienceAPI, 0)
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

// Update experience
// PUT /v1/experiences/:id
// https://connectustoday.github.io/api-server/api-reference#experiences

func UpdateExperienceRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params, account interface{}) {
	//TODO REDO THE CODE (MULTIPLE SAVES IN PARALLEL), SAVE AND TRANSFER APPROVAL FROM ORGANIZATION IF ORGANIZATION REMAINS THE SAME
	w.Header().Set("Content-Type", "application/json")

}

// Create experience
// POST /v1/experiences
// https://connectustoday.github.io/api-server/api-reference#experiences

func CreateExperienceRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params, account interface{}) {
	w.Header().Set("Content-Type", "application/json")

	type requestForm struct {
		Location     *interfaces_api.IAddressAPI `json:"location" schema:"location"`
		Name         *string                     `json:"name" schema:"name"`
		Organization *string                     `json:"organization" schema:"organization"`
		Opportunity  *string                     `json:"opportunity" schema:"opportunity"`
		Description  *string                     `json:"description" schema:"description"`
		When         *interfaces_api.When        `json:"when" schema:"when"`
		Hours        *int                        `json:"hours" schema:"hours"`
		EmailVerify  *bool                       `json:"email_verify" schema:"email_verify"`
	}

	var req requestForm

	err := DecodeRequest(r, &req)
	if err != nil { // Check decoding error
		SendError(w, http.StatusInternalServerError, internalServerError, 4001)
		return
	}
	if !VerifyFieldsExist(req, FormOmit([]string{"Location", "Organization", "Opportunity", "When", "Hours"}), true) { // Check request for correct fields
		SendError(w, http.StatusBadRequest, badRequest+" (Bad request.)", 4001) // TODO proper error code
		return
	}

	user, ok := account.(interfaces_internal.IUser)
	if !ok || user.Type != "User" { // check if the obtained account is of user type
		SendError(w, http.StatusBadRequest, badRequest+"  (Incorrect account type! User account type required.)", 4000)
	}

	// verifications for data

	if *req.Opportunity != "" {
		// TODO opportunity
	}

	if *req.Organization != "" { // if the organization field is filled out
		if *req.EmailVerify { // send email request to organization not on site (for validations)

		} else {
			var org interfaces_internal.IOrganization
			err := IAccountCollection.Find(bson.M{"username": *req.Organization, "type": "Organization"}).One(&org) // TODO CASE INSENSITIVE LOOKUPS
			if err != nil {
				if err.Error() == "not found" {
					SendError(w, http.StatusBadRequest, badRequest + " (Organization not found.)", 4002)
				} else {
					SendError(w, http.StatusInternalServerError, internalServerError, 4001)
				}
				return
			}

			// TODO
			org.ExperienceValidations = append(org.ExperienceValidations, interfaces_internal.IValidations{
				UserID:       user.UserName,
				ExperienceID: re, // TODO BSON OBJECT ID
			})
		})
	}



}

// Delete experience
// DELETE /v1/experiences/:id
// https://connectustoday.github.io/api-server/api-reference#experiences

func DeleteExperienceRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params, account interface{}) {
	w.Header().Set("Content-Type", "application/json")

}

// View experience validations
// GET /v1/experiences/validations
// https://connectustoday.github.io/api-server/api-reference#experiences

func GetExperienceValidationsRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params, account interface{}) {
	w.Header().Set("Content-Type", "application/json")

}

// Approve or don't approve an experience validation request
// POST /v1/experiences/validations/:user/:id
// https://connectustoday.github.io/api-server/api-reference#experiences

func ReviewExperienceValidationsRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params, account interface{}) {
	w.Header().Set("Content-Type", "application/json")

}

// Approve Validation (From email instead of account)
// POST /v1/experiences/email_approve/:token
// https://connectustoday.github.io/api-server/api-reference#experiences

func EmailApproveExperienceValidationRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params, account interface{}) {
	w.Header().Set("Content-Type", "application/json")

}
