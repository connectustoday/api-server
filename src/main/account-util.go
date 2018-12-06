package main

import (
	"encoding/json"
	"github.com/globalsign/mgo/bson"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	"interfaces-api"
	"interfaces-internal"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

func VerifyUniqueUsername(username string) bool {
	count, err := IAccountCollection.Find(bson.M{"username": username}).Count()
	if err != nil {
		log.Print(err)
		return false
	}
	return count == 0
}

/*
 * Account registration route
 * POST /v1/auth/register
 * https://connectustoday.github.io/api-server/api-reference#register
 */

func RegisterRoute(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	type requestForm struct {
		// Global account fields
		UserName *string `json:"username,omitempty"`
		Email    *string `json:"email"`
		Password *string `json:"password"`
		Type     *string `json:"type"`

		// User specific fields
		FirstName *string `json:"first_name"`
		Birthday  *string `json:"string"`

		// Organization specific fields
		IsNonProfit   *bool   `json:"is_nonprofit"`
		PreferredName *string `json:"preferred_name"`
	}

	var req requestForm

	//body, _ := ioutil.ReadAll(r.Body)
	//r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	// TODO SWITCH TO GENERAL UTIL FUNCTION FOR DECODING

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil { // Check decoding error
		if DEBUG {
			output, _ := httputil.DumpRequest(r, true)
			println(string(output))
			log.Println(err.Error() + " ")
		}
		SendError(w, http.StatusInternalServerError, internalServerError+" (There was a problem reading the request.)", 3205)
		return
	}
	if !VerifyFieldsExist(req, FormOmit([]string{"FirstName", "Birthday", "IsNonProfit", "PreferredName"})) { // Check request for correct fields
		SendError(w, http.StatusBadRequest, badRequest+" (Bad request.)", 3206)
		return
	}
	if !VerifyUniqueUsername(*req.UserName) { // Check if username is unique
		SendError(w, http.StatusBadRequest, badRequest+" (Username already taken.)", 3201)
		return
	}

	// Get hashed bcrypt password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*req.Password), 14)
	if err != nil { // check for bcrypt error
		SendError(w, http.StatusInternalServerError, internalServerError+" (There was a problem registering the account.)", 3203)
		return
	}

	// Send email verification email
	// TODO ACTUALLY DO EMAIL

	if *req.Type == "user" {

		err = IAccountCollection.Insert(interfaces_internal.IUser{ // Add Default User
			IAccount: &interfaces_internal.IAccount{
				SchemaVersion:        0,
				UserName:             *req.UserName,
				Email:                *req.Email,
				Password:             string(hashedPassword), // TODO verify if this is how it works or [:]
				OAuthToken:           "",
				OAuthService:         "",
				IsEmailVerified:      false,
				LastLogin:            0,
				Notifications:        []interfaces_internal.INotification{},
				Avatar:               "https://pbs.twimg.com/profile_images/1017516299143041024/fLFdcGsl_400x400.jpg", // TODO default images
				Header:               "https://pbs.twimg.com/profile_images/1017516299143041024/fLFdcGsl_400x400.jpg",
				CreatedAt:            time.Now().Unix(),
				PendingConnections:   []string{},
				RequestedConnections: []string{},
				Posts:                []string{},
				Liked:                []interfaces_internal.ICom{},
				Shared:               []interfaces_internal.ICom{},
				Settings: interfaces_internal.IUserSettings{
					IAccountSettings: &interfaces_internal.IAccountSettings{
						AllowMessagesFromUnknown: true,
						EmailNotifications:       false,
					},
					IsFullNameVisible: true,
					BlockedUsers:      []string{},
				},
				AdminNote: "",
				Type:      "user",
			},
			FirstName:  *req.FirstName,
			MiddleName: "",
			LastName:   "",
			Birthday:   *req.Birthday,
			Gender:     "",
			PersonalInfo: interfaces_internal.IUserProfile{
				SchemaVersion:    0,
				Interests:        []string{},
				Biography:        "",
				Education:        "",
				Quote:            "",
				CurrentResidence: "",
				Certifications:   "",
			},
			Experiences: []interfaces_internal.IExperience{},
		})

		// Check for successful insert to database

		if err != nil {
			SendError(w, http.StatusInternalServerError, internalServerError+" (There was a problem registering the account.)", 3203)
		} else {
			err = WriteOK(w)
			if err != nil {
				println(err)
			}
		}
	} else if *req.Type == "organization" {

		err = IAccountCollection.Insert(interfaces_internal.IOrganization{ // Add Default Organization
			IAccount: &interfaces_internal.IAccount{
				SchemaVersion:        0,
				UserName:             *req.UserName,
				Email:                *req.Email,
				Password:             string(hashedPassword),
				OAuthToken:           "",
				OAuthService:         "",
				IsEmailVerified:      false,
				LastLogin:            0,
				Notifications:        []interfaces_internal.INotification{},
				Avatar:               "https://pbs.twimg.com/profile_images/1017516299143041024/fLFdcGsl_400x400.jpg",
				Header:               "https://pbs.twimg.com/profile_images/1017516299143041024/fLFdcGsl_400x400.jpg",
				CreatedAt:            time.Now().Unix(),
				PendingConnections:   []string{},
				RequestedConnections: []string{},
				Posts:                []string{},
				Liked:                []interfaces_internal.ICom{},
				Shared:               []interfaces_internal.ICom{},
				Settings:             interfaces_internal.IOrganizationSettings{
					IAccountSettings: &interfaces_internal.IAccountSettings{
						AllowMessagesFromUnknown: true,
						EmailNotifications:       true,
					},
					IsNonprofit: *req.IsNonProfit,
				},
				AdminNote:            "",
				Type:                 "",
			},
			PreferredName: *req.PreferredName,
			IsVerified:    false,
			Opportunities: []string{},
			OrgInfo: interfaces_internal.IOrganizationProfile{
				SchemaVersion: 0,
				Mission:       "",
				Quote:         "",
				Address: interfaces_internal.IAddress{
					SchemaVersion: 0,
					Street:        "",
					City:          "",
					Province:      "",
					Country:       "",
					PostalCode:    "",
					AptNumber:     "",
					GeoJSON: interfaces_internal.IPoint{
						Type:        "",
						Coordinates: nil,
					},
				},
				AffiliatedOrgs: []string{},
				Interests:      []string{},
			},
			ExperienceValidations: []interfaces_internal.IValidations{},
		})

		// Check for successful insert to database
		if err != nil {
			SendError(w, http.StatusInternalServerError, internalServerError+" (There was a problem registering the account.)", 3203)
		} else {
			err = WriteOK(w)
			if err != nil {
				println(err)
			}
		}
	} else {
		SendError(w, http.StatusBadRequest, badRequest+" (Invalid account type)", 3200)
	}
}

func VerifyEmailRequestRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

/*
 * Get account route
 * GET /v1/accounts/:id
 * https://connectustoday.github.io/api-server/api-reference#accounts
 */

func GetAccountRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	account := interfaces_internal.IAccount{}
	err := IAccountCollection.Find(bson.M{"username": p.ByName("id")}).One(&account)

	if err != nil {
		if err.Error() == "not found" {
			SendError(w, http.StatusNotFound, notFound+" (Account not found)", 3003)
		} else {
			SendError(w, http.StatusInternalServerError, internalServerError+" (Problem finding account)", 3002)
		}
		return
	}

	accountapi := interfaces_api.ConvertToIAccountAPI(account)

	b, err := json.Marshal(accountapi)
	if err != nil {
		SendError(w, http.StatusInternalServerError, internalServerError+" (Problem finding account)", 3002)
		return
	}

	_, err = w.Write([]byte(b))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func GetAccountProfileRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func GetAccountConnectionsRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
