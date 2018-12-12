package main

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/globalsign/mgo/bson"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"interfaces-api"
	"interfaces-conv"
	"interfaces-internal"
	"log"
	"mail-templates"
	"net/http"
	"strconv"
	"time"
)

// Get experiences of personal user (username found from token)
// GET /v1/experiences
// https://connectustoday.github.io/api-server/api-reference#experiences

func GetPersonalExperiencesRoute(w http.ResponseWriter, _ *http.Request, _ httprouter.Params, account bson.M) {
	w.Header().Set("Content-Type", "application/json")

	user, err := interfaces_conv.ConvertBSONToIUser(account)
	if err != nil || user.Type != "User" { // check if the obtained account is of user type
		SendError(w, http.StatusBadRequest, badRequest+"  (Incorrect account type! User account type required.)", 4000)
		return
	}

	GetExperiences(w, user)
}

// Get all experiences of any user
// GET /v1/accounts/:id/experiences
// https://connectustoday.github.io/api-server/api-reference#accounts

func GetExperiencesRoute(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	var user interfaces_internal.IUser

	// Find and get user object
	err := IAccountCollection.Find(bson.M{"username": p.ByName("id"), "type": "User"}).One(&user)
	if CheckMongoQueryError(w, err, "  (User not found! Is this the correct account type?)", 4002, 4001) != nil {
		return
	}

	GetExperiences(w, user)
}

// Get experience helper

func GetExperiences(w http.ResponseWriter, user interfaces_internal.IUser) {
	ret := make([]interfaces_api.IExperienceAPI, 0)
	for _, e := range user.Experiences { // add experiences to api array
		ret = append(ret, interfaces_conv.ConvertToIExperienceAPI(e))
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

func UpdateExperienceRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params, account bson.M) {
	// TODO REDO THE CODE (MULTIPLE SAVES IN PARALLEL), SAVE AND TRANSFER APPROVAL FROM ORGANIZATION IF ORGANIZATION REMAINS THE SAME
	w.Header().Set("Content-Type", "application/json")

	// TODO
}

// Create experience
// POST /v1/experiences
// https://connectustoday.github.io/api-server/api-reference#experiences

func CreateExperienceRoute(w http.ResponseWriter, r *http.Request, _ httprouter.Params, account bson.M) {
	w.Header().Set("Content-Type", "application/json")

	type requestForm struct {
		Location     *interfaces_api.IAddressAPI `json:"location" schema:"location"`
		Name         *string                     `json:"name" schema:"name"`
		Organization *string                     `json:"organization" schema:"organization"`
		Opportunity  *string                     `json:"opportunity" schema:"opportunity"`
		Description  *string                     `json:"description" schema:"description"`
		When         *interfaces_internal.When   `json:"when" schema:"when"`
		Hours        *int64                      `json:"hours" schema:"hours"`
		EmailVerify  *bool                       `json:"email_verify" schema:"email_verify"`
	}

	var req requestForm

	err := DecodeRequest(r, &req)
	if err != nil { // Check decoding error
		if DEBUG {
			log.Println(err)
		}
		SendError(w, http.StatusInternalServerError, internalServerError+" when parsing request. Bad input?", 4050)
		return
	}
	if !VerifyFieldsExist(&req, FormOmit([]string{"Location", "Organization", "Opportunity", "When", "Hours"}), true) { // Check request for correct fields
		SendError(w, http.StatusBadRequest, badRequest+" (Bad request.)", 4050) // TODO proper error code
		return
	}

	user, err := interfaces_conv.ConvertBSONToIUser(account)
	if err != nil || user.Type != "User" { // check if the obtained account is of user type
		SendError(w, http.StatusBadRequest, badRequest+"  (Incorrect account type! User account type required.)", 4000)
		return
	}

	// convert to internal experience

	exp := interfaces_internal.IExperience{
		ID:            bson.NewObjectId(),
		SchemaVersion: 0,
		Location:      interfaces_conv.ConvertToIAddressInternal(*req.Location),
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
		if exp.EmailVerify { // send email request to organization not on site (for validations)
			token := jwt.New(jwt.SigningMethodHS256)

			// create jwt token for organization verification
			claims := make(jwt.MapClaims)
			claims["username"] = user.UserName
			claims["ms"] = time.Now().Unix()
			claims["exp"] = time.Now().Add(time.Second * time.Duration(604800)).Unix() // expires in one week
			token.Claims = claims
			tokenString, err := token.SignedString([]byte(APPROVAL_VERIFY_SECRET)) // sign with secret
			if err != nil {
				SendError(w, http.StatusInternalServerError, internalServerError, 4001)
				return
			}

			exp.EmailJWT = tokenString
			verifyLink := API_DOMAIN + "/v1/experiences/email-approve/" + exp.EmailJWT

			err = SendMail(exp.Organization, "Volunteer or Work Experience Validation Request", mail_templates.VALIDATE_EXPERIENCE_WITHOUT_ACCOUNT, struct {
				VerifyLink template.URL
				Website    template.URL
				Email      string
				UserName   string
				FullName   string
				ExpName    string
				ExpHours   string
				ExpStart   string
				ExpEnd     string
				ExpDesc    string
			}{template.URL(verifyLink),
				template.URL(SITE_DOMAIN),
				user.Email,
				user.UserName,
				user.FirstName + " " + user.MiddleName + " " + user.LastName,
				exp.Name,
				strconv.Itoa(int(exp.Hours)),
				exp.When.Begin,
				exp.When.End,
				exp.Description})

			if err != nil {
				if DEBUG {
					log.Println(err)
				}
				SendError(w, http.StatusInternalServerError, internalServerError+" (Issue sending mail)", 4003)
				return
			}

		} else { // check if there is an associated organization on the site (for validations)

			// add validation request to organization pending validations list
			// TODO NOTIFICATION ON PENDING VALIDATION
			// TODO DUPLICATE HEADERS SENT WHEN SAVING FAILURE

			var org interfaces_internal.IOrganization
			err := IAccountCollection.Find(bson.M{"username": exp.Organization, "type": "Organization"}).One(&org) // TODO CASE INSENSITIVE LOOKUPS
			if CheckMongoQueryError(w, err, " (Organization not found.)", 4002, 4001) != nil {
				return
			}

			org.ExperienceValidations = append(org.ExperienceValidations, interfaces_internal.IValidations{
				UserID:       user.UserName,
				ExperienceID: exp.ID.Hex(), // TODO BSON OBJECT ID
			}) // add validation entry to organization

			err = IAccountCollection.Update(bson.M{"username": exp.Organization, "type": "Organization"}, org) // save to db
			if CheckMongoQueryError(w, err, internalServerError, 4001, 4001) != nil {
				return
			}
		}
	}

	user.Experiences = append(user.Experiences, exp) // add to user's experiences array
	// finish adding experience to database
	err = IAccountCollection.Update(bson.M{"username": user.UserName}, user)
	if err != nil {
		SendError(w, http.StatusInternalServerError, internalServerError, 4001)
		return
	}

	WriteOK(w)
}

// Delete experience
// DELETE /v1/experiences/:id
// https://connectustoday.github.io/api-server/api-reference#experiences

func DeleteExperienceRoute(w http.ResponseWriter, _ *http.Request, p httprouter.Params, account bson.M) {
	w.Header().Set("Content-Type", "application/json")

	// check if the obtained account is of user type and convert
	user, err := interfaces_conv.ConvertBSONToIUser(account)
	if err != nil || user.Type != "User" {
		SendError(w, http.StatusBadRequest, badRequest+"  (Incorrect account type! User account type required.)", 4000)
		return
	}

	var exp interfaces_internal.IExperience
	expIndex := -1

	// find experience to delete in array

	for i, ex := range user.Experiences {
		if ex.ID.Hex() == p.ByName("id") {
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
		found := true
		err := IAccountCollection.Find(bson.M{"username": exp.Organization, "type": "Organization"}).One(&org) // TODO CASE INSENSITIVE LOOKUPS

		if err != nil {
			if err.Error() == "not found" {
				found = false
			} else {
				SendError(w, http.StatusInternalServerError, internalServerError, 4001)
				return
			}
		}

		// remove all entries with the same id and user (duplicates as well) if there is an organization with a request
		if found {
			found = false
			i := 0

			// remove all entries with the same id and user (duplicates as well)
			for _, v := range org.ExperienceValidations {
				if v.UserID == user.UserName && v.ExperienceID == exp.ID.Hex() {
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
	err = IAccountCollection.Update(bson.M{"username": user.UserName}, user)
	if err != nil {
		SendError(w, http.StatusInternalServerError, internalServerError, 4001)
		return
	}
	WriteOK(w)
}

// View experience validations
// GET /v1/experiences/validations
// https://connectustoday.github.io/api-server/api-reference#experiences

func GetExperienceValidationsRoute(w http.ResponseWriter, _ *http.Request, _ httprouter.Params, account bson.M) {
	w.Header().Set("Content-Type", "application/json")

	// check if the obtained account is of organization type and convert
	org, err := interfaces_conv.ConvertBSONToIOrganization(account)
	if err != nil || org.Type != "Organization" {
		SendError(w, http.StatusBadRequest, badRequest+"  (Incorrect account type! Organization account type required.)", 4000)
		return
	}

	ret := make([]interfaces_api.IValidationsAPI, 0)
	for _, e := range org.ExperienceValidations { // add experiences to api array
		ret = append(ret, interfaces_conv.ConvertToIValidationsAPI(e))
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

func ReviewExperienceValidationsRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params, account bson.M) {
	w.Header().Set("Content-Type", "application/json")

	type requestForm struct {
		Approve *bool `json:"approve" schema:"approve"`
	}

	var req requestForm

	err := DecodeRequest(r, &req)
	if err != nil { // Check decoding error
		SendError(w, http.StatusInternalServerError, internalServerError+" when parsing request. Bad input?", 4050)
		return
	}
	if !VerifyFieldsExist(&req, FormOmit([]string{}), true) { // Check request for correct fields
		SendError(w, http.StatusBadRequest, badRequest+" (Bad request.)", 4050) // TODO proper error code
		return
	}

	// check if the obtained account is of organization type and convert
	org, err := interfaces_conv.ConvertBSONToIOrganization(account)
	if err != nil || org.Type != "Organization" {
		SendError(w, http.StatusBadRequest, badRequest+"  (Incorrect account type! Organization account type required.)", 4000)
		return
	}

	found := false

	// Remove the experience validation request from the organization object
	for i := 0; i < len(org.ExperienceValidations); i++ {
		if org.ExperienceValidations[i].UserID == p.ByName("user") && org.ExperienceValidations[i].ExperienceID == p.ByName("id") {
			found = true
			org.ExperienceValidations = append(org.ExperienceValidations[:i], org.ExperienceValidations[i+1:]...)
			i--
		}
		i++
	}

	if !found {
		SendError(w, http.StatusNotFound, notFound+" (Experience validation not found)", 4002)
		return
	}

	// Save organization to db
	err = IAccountCollection.Update(bson.M{"username": org.UserName}, org)
	if CheckMongoQueryError(w, err, internalServerError, 4001, 4001) != nil {
		return
	}

	// Update user experience object with approval
	var user interfaces_internal.IUser
	err = IAccountCollection.Find(bson.M{"username": p.ByName("user"), "type": "User"}).One(&user)
	if CheckMongoQueryError(w, err, badRequest+" (User not found.)", 4003, 4001) != nil {
		return
	}

	found = false

	// Update user's experience
	for i, ex := range user.Experiences {
		if ex.ID.Hex() == p.ByName("id") {
			if *req.Approve {
				user.Experiences[i].IsVerified = true // verify experience object if approved
			} else {
				user.Experiences = append(user.Experiences[:i], user.Experiences[i+1:]...) // delete experience object if not approved
			}
			found = true
			break
		}
	}

	if !found {
		SendError(w, http.StatusBadRequest, internalServerError+" (Experience not found in user object)", 4004)
		return
	}

	err = IAccountCollection.Update(bson.M{"username": p.ByName("user"), "type": "User"}, user)
	if CheckMongoQueryError(w, err, internalServerError, 4001, 4001) != nil {
		return
	}

	WriteOK(w)
}

// Approve Validation (From email instead of account)
// POST /v1/experiences/email_approve/:token
// https://connectustoday.github.io/api-server/api-reference#experiences

func EmailApproveExperienceValidationRoute(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")

	claims, err := GetJWTClaims(p.ByName("token"), APPROVAL_VERIFY_SECRET) // verify token authenticity

	if err != nil {
		w.WriteHeader(404)
		_, _ = w.Write([]byte("Invalid approval link. Perhaps it has expired?"))
		return
	}

	var user interfaces_internal.IUser
	err = IAccountCollection.Find(bson.M{"username": claims["username"]}).One(&user)
	if err != nil {
		if err.Error() == "not found" {
			w.WriteHeader(404)
			_, _ = w.Write([]byte("Account not found. Perhaps the user has been removed?"))
		} else {
			w.WriteHeader(500)
			_, _ = w.Write([]byte("Internal server error. Problem finding account."))
		}
		return
	}

	found := false
	for i := range user.Experiences {
		if user.Experiences[i].EmailJWT != "" {
			claims2, _ := GetJWTClaims(user.Experiences[i].EmailJWT, APPROVAL_VERIFY_SECRET)
			if claims2["ms"] == claims["ms"] {
				user.Experiences[i].EmailJWT = ""
				user.Experiences[i].IsVerified = true
				found = true
			}
		}
	}

	if !found {
		w.WriteHeader(500)
		_, _ = w.Write([]byte("Could not find experience to validate. Perhaps the user has removed it, or the experience was already validated?"))
		return
	}

	err = IAccountCollection.Update(bson.M{"username": claims["username"]}, user)

	if err != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte("Internal server error. :("))
		return
	}

	_, _ = w.Write([]byte("You have successfully approved the request! Sign up for ConnectUS to approve and manage validations directly from the site...<script>setTimeout(()=>{window.location.replace('" + SITE_DOMAIN + "/auth/login.php')}, 5000)</script>"))
}
