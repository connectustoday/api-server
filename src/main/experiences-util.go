package main

import (
	"encoding/json"
	"github.com/globalsign/mgo/bson"
	"github.com/julienschmidt/httprouter"
	"interfaces-api"
	"interfaces-internal"
	"net/http"
	"time"
)

// Get experiences of personal user (username found from token)
// GET /v1/experiences
// https://connectustoday.github.io/api-server/api-reference#experiences

func GetPersonalExperiencesRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params, account interface{}) {
	w.Header().Set("Content-Type", "application/json")

	user, ok := account.(interfaces_internal.IUser)
	if !ok || user.Type != "User" { // check if the obtained account is of user type
		SendError(w, http.StatusBadRequest, badRequest+"  (Incorrect account type! User account type required.)", 4000)
		return
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
	if checkMongoQueryError(err, "  (User not found! Is this the correct account type?)", 4002, 4001) != nil {
		return
	}

	GetExperiences(w, user)
}

// Get experience helper

func GetExperiences(w http.ResponseWriter, user interfaces_internal.IUser) {
	ret := make([]interfaces_api.IExperienceAPI, 0)
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
		Location     *interfaces_internal.IAddress `json:"location" schema:"location"`
		Name         *string                       `json:"name" schema:"name"`
		Organization *string                       `json:"organization" schema:"organization"`
		Opportunity  *string                       `json:"opportunity" schema:"opportunity"`
		Description  *string                       `json:"description" schema:"description"`
		When         *interfaces_internal.When     `json:"when" schema:"when"`
		Hours        *int64                        `json:"hours" schema:"hours"`
		EmailVerify  *bool                         `json:"email_verify" schema:"email_verify"`
	}

	var req requestForm

	err := DecodeRequest(r, &req)
	if err != nil { // Check decoding error
		SendError(w, http.StatusInternalServerError, internalServerError, 4001)
		return
	}
	if !VerifyFieldsExist(req, FormOmit([]string{"Location", "Organization", "Opportunity", "When", "Hours"}), true) { // Check request for correct fields
		SendError(w, http.StatusBadRequest, badRequest+" (Bad request.)", 4004) // TODO proper error code
		return
	}

	user, ok := account.(interfaces_internal.IUser)
	if !ok || user.Type != "User" { // check if the obtained account is of user type
		SendError(w, http.StatusBadRequest, badRequest+"  (Incorrect account type! User account type required.)", 4000)
		return
	}

	// convert to experience

	exp := interfaces_internal.IExperience{
		ID:            bson.NewObjectId(),
		SchemaVersion: 0,
		Location:      *req.Location,
		Name:          *req.Name,
		Organization:  *req.Organization,
		Opportunity:   *req.Opportunity,
		Description:   *req.Description,
		When:          *req.When,
		IsVerified:    false,
		EmailVerify:   *req.EmailVerify,
		CreatedAt:     time.Now().Unix(),
		Hours:         *req.Hours,
		EmailJWT:      "",
	}

	// verifications for data

	if exp.Opportunity != "" {
		// TODO opportunity
	}

	if exp.Organization != "" { // if the organization field is filled out
		if *req.EmailVerify { // send email request to organization not on site (for validations)

		} else {
			var org interfaces_internal.IOrganization
			err := IAccountCollection.Find(bson.M{"username": exp.Organization, "type": "Organization"}).One(&org) // TODO CASE INSENSITIVE LOOKUPS
			if checkMongoQueryError(err, " (Organization not found.)", 4002, 4001) != nil {
				return
			}

			// TODO
			org.ExperienceValidations = append(org.ExperienceValidations, interfaces_internal.IValidations{
				UserID:       user.UserName,
				ExperienceID: re, // TODO BSON OBJECT ID
			})
		}
	}

	user.Experiences = append(user.Experiences, exp) // add to user's experiences array
	// finish adding experience to database
	err = IAccountCollection.Update(bson.M{"username": user.UserName}, user)
	if err != nil {
		SendError(w, http.StatusInternalServerError, internalServerError, 4001)
		return
	}
	err = WriteOK(w)
	if err != nil {
		SendError(w, http.StatusInternalServerError, internalServerError, 4001)
		return
	}
}

// Delete experience
// DELETE /v1/experiences/:id
// https://connectustoday.github.io/api-server/api-reference#experiences

func DeleteExperienceRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params, account interface{}) {
	w.Header().Set("Content-Type", "application/json")

	// check if the obtained account is of user type and convert
	user, ok := account.(interfaces_internal.IUser)
	if !ok || user.Type != "User" {
		SendError(w, http.StatusBadRequest, badRequest+"  (Incorrect account type! User account type required.)", 4000)
		return
	}

	var exp interfaces_internal.IExperience
	expIndex := -1

	// find experience to delete in array
	for i, ex := range user.Experiences {
		if ex.ID.String() == p.ByName("id") {
			exp = ex
			expIndex = i
			break
		}
	}

	if expIndex < 0 { // if not found
		SendError(w, http.StatusNotFound, notFound+" (Experience not found with supplied ID)", 4002)
		return
	}

	if exp.Opportunity != "" {
		// TODO OPPORTUNITY
	}

	if exp.Organization != "" && !exp.EmailVerify { // remove pending validations for experience from organization
		var org interfaces_internal.IOrganization
		found := false
		err := IAccountCollection.Find(bson.M{"username": exp.Organization, "type": "Organization"}).One(&org) // TODO CASE INSENSITIVE LOOKUPS

		if err != nil {
			if err.Error() == "not found" {
				found = true
			} else {
				SendError(w, http.StatusInternalServerError, internalServerError, 4001)
				return
			}
		}

		if found { // remove all entries with the same id and user (duplicates as well)
			found = false
			i := 0

			for _, v := range org.ExperienceValidations { // remove all entries with the same id and user (duplicates as well)
				if v.UserID == user.UserName && v.ExperienceID == exp.ID.String() {
					found = true
					org.ExperienceValidations = append(org.ExperienceValidations[:i], org.ExperienceValidations[i+1:]...) // remove from slice
					i--
				}
				i++
			}

			if found { // if it exists
				err := IAccountCollection.Update(bson.M{"username": exp.Organization, "type": "Organization"}, org)
				if err != nil {
					SendError(w, http.StatusInternalServerError, internalServerError, 4001)
					return
				}
			}
		}
	}

	// Save user object
	user.Experiences = append(user.Experiences[:expIndex], user.Experiences[expIndex+1:]...)
	err := IAccountCollection.Update(bson.M{"username": user.UserName}, user)
	if err != nil {
		SendError(w, http.StatusInternalServerError, internalServerError, 4001)
		return
	}
	err = WriteOK(w)
	if err != nil {
		SendError(w, http.StatusInternalServerError, internalServerError, 4001)
		return
	}
}

// View experience validations
// GET /v1/experiences/validations
// https://connectustoday.github.io/api-server/api-reference#experiences

func GetExperienceValidationsRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params, account interface{}) {
	w.Header().Set("Content-Type", "application/json")

	// check if the obtained account is of organization type and convert
	org, ok := account.(interfaces_internal.IOrganization)
	if !ok || org.Type != "Organization" {
		SendError(w, http.StatusBadRequest, badRequest+"  (Incorrect account type! Organization account type required.)", 4000)
		return
	}

	ret := make([]interfaces_api.IValidationsAPI, 0)
	for _, e := range org.ExperienceValidations { // add experiences to api array
		ret = append(ret, interfaces_api.ConvertToIValidationsAPI(e))
	}

	response, err := json.Marshal(ret) // prepare json response
	if err != nil {
		SendError(w, http.StatusInternalServerError, internalServerError, 4001)
		return
	}

	_, err = w.Write(response) // write response to client
	if err != nil {
		SendError(w, http.StatusInternalServerError, internalServerError, 4001)
		return
	}
}

// Approve or don't approve an experience validation request
// POST /v1/experiences/validations/:user/:id
// https://connectustoday.github.io/api-server/api-reference#experiences

func ReviewExperienceValidationsRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params, account interface{}) {
	w.Header().Set("Content-Type", "application/json")

	type requestForm struct {
		Approve bool `json:"approve" schema:"approve"`
	}

	var req requestForm

	err := DecodeRequest(r, &req)
	if err != nil { // Check decoding error
		SendError(w, http.StatusInternalServerError, internalServerError, 4001)
		return
	}
	if !VerifyFieldsExist(req, FormOmit([]string{}), true) { // Check request for correct fields
		SendError(w, http.StatusBadRequest, badRequest+" (Bad request.)", 4004) // TODO proper error code
		return
	}

	// check if the obtained account is of organization type and convert
	org, ok := account.(interfaces_internal.IOrganization)
	if !ok || org.Type != "Organization" {
		SendError(w, http.StatusBadRequest, badRequest+"  (Incorrect account type! Organization account type required.)", 4000)
		return
	}

	foundIndex := -1

	for i, v := range org.ExperienceValidations {
		if v.UserID == p.ByName("user") && v.ExperienceID == p.ByName("id") {
			foundIndex = 
		}
	}
}

// Approve Validation (From email instead of account)
// POST /v1/experiences/email_approve/:token
// https://connectustoday.github.io/api-server/api-reference#experiences

func EmailApproveExperienceValidationRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params, account interface{}) {
	w.Header().Set("Content-Type", "application/json")

}

func checkMongoQueryError(err error, notFoundMsg string, errCodeNotFound int, errCodeError int) error {
	if err != nil {
		if err.Error() == "not found" {
			SendError(w, http.StatusNotFound, notFound+notFoundMsg, errCodeNotFound)
		} else {
			SendError(w, http.StatusInternalServerError, internalServerError, errCodeError)
		}
	}
	return err
}