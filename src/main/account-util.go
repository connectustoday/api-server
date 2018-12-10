package main

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/globalsign/mgo/bson"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"interfaces-conv"
	"interfaces-internal"
	"log"
	"mail-templates"
	"net/http"
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
	w.Header().Set("Content-Type", "application/json")

	type requestForm struct {
		// Global account fields
		UserName *string `json:"username" schema:"username"`
		Email    *string `json:"email" schema:"email"`
		Password *string `json:"password" schema:"password"`
		Type     *string `json:"type" schema:"type"`

		// User specific fields
		FirstName *string `json:"first_name" schema:"first_name"`
		Birthday  *string `json:"string" schema:"birthday"`

		// Organization specific fields
		IsNonProfit   *bool   `json:"is_nonprofit" schema:"is_non_profit"`
		PreferredName *string `json:"preferred_name" schema:"preferred_name"`
	}

	var req requestForm

	err := DecodeRequest(r, &req)
	if err != nil { // Check decoding error
		if DEBUG {
			//output, _ := httputil.DumpRequest(r, true)
			//println(string(output))
			log.Println(err.Error())
		}
		SendError(w, http.StatusInternalServerError, internalServerError+" (There was a problem reading the request.)", 3205)
		return
	}
	if !VerifyFieldsExist(&req, FormOmit([]string{"FirstName", "Birthday", "IsNonProfit", "PreferredName"}), false) { // Check request for correct fields
		SendError(w, http.StatusBadRequest, badRequest+" (Bad request.)", 3206)
		return
	}
	if !VerifyUniqueUsername(*req.UserName) { // Check if username is unique
		SendError(w, http.StatusBadRequest, badRequest+" (Username already taken.)", 3201)
		return
	}

	// Get hashed bcrypt password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*req.Password), BCRYPT_COST)
	if err != nil { // check for bcrypt error
		SendError(w, http.StatusInternalServerError, internalServerError+" (There was a problem registering the account.)", 3203)
		return
	}

	// Send email verification email
	if err = sendVerificationEmail(*req.UserName, *req.Email); err != nil {
		log.Printf("Problem sending mail or creating verification token: %s", err.Error())
		SendError(w, http.StatusInternalServerError, internalServerError+" (There was a problem sending the verification email. Please ask a website administrator for help.)", 3204)
		return
	}

	if *req.Type == "user" {

		if !VerifyFieldsExist(&req, FormOmit([]string{"FirstName", "Birthday", "IsNonProfit", "PreferredName"}), true) { // Check request for correct fields
			SendError(w, http.StatusBadRequest, badRequest+" (Bad request.)", 3206)
			return
		}

		err = IAccountCollection.Insert(&interfaces_internal.IUser{ // Add Default User
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
				AllowMessagesFromUnknown: true,
				EmailNotifications:       false,
				IsFullNameVisible:        true,
				BlockedUsers:             []string{},
			},
			AdminNote: "",
			Type:      "User",
			// user specific fields
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
			WriteOK(w)
		}
	} else if *req.Type == "organization" {

		if !VerifyFieldsExist(&req, FormOmit([]string{"FirstName", "Birthday", "IsNonProfit", "PreferredName"}), true) { // Check request for correct fields
			SendError(w, http.StatusBadRequest, badRequest+" (Bad request.)", 3206)
			return
		}

		err = IAccountCollection.Insert(&interfaces_internal.IOrganization{ // Add Default Organization
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
			Settings: interfaces_internal.IOrganizationSettings{
				AllowMessagesFromUnknown: true,
				EmailNotifications:       true,
				IsNonprofit:              *req.IsNonProfit,
			},
			AdminNote: "",
			Type:      "Organization",
			// organization specific fields
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
			WriteOK(w)
		}
	} else {
		SendError(w, http.StatusBadRequest, badRequest+" (Invalid account type)", 3200)
	}
}

func sendVerificationEmail(username string, email string) error {
	token := jwt.New(jwt.SigningMethodHS256)

	// create jwt token for organization verification
	claims := make(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Second * time.Duration(43200)).Unix() // expires in one week
	token.Claims = claims
	tokenString, err := token.SignedString([]byte(REGISTER_VERIFY_SECRET)) // sign with secret
	if err != nil {
		return err
	}
	verifyLink := API_DOMAIN + "/v1/auth/verify-email/" + tokenString
	return SendMail(email, "ConnectUS Account Verification Code", mail_templates.REGISTER_VERIFY, struct {
		VerifyLink template.URL
	}{VerifyLink: template.URL(verifyLink)})
}

// Verify email request route
// https://connectus.github.io/api-server/api-reference#accounts

func VerifyEmailRequestRoute(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")

	claims, err := GetJWTClaims(p.ByName("token"), REGISTER_VERIFY_SECRET) // verify token authenticity
	if err != nil {
		w.WriteHeader(404)
		_, _ = w.Write([]byte("Invalid verification link. Perhaps it's expired?"))
		return
	}

	var acc interfaces_internal.IAccount

	err = IAccountCollection.Find(bson.M{"username": claims["username"]}).One(&acc)

	if err != nil {
		if err.Error() == "not found" {
			w.WriteHeader(500)
			_, _ = w.Write([]byte("Account not found. Please try registering again."))
		} else {
			w.WriteHeader(500)
			_, _ = w.Write([]byte("Internal server error. Problem finding account."))
		}
		return
	}

	if acc.IsEmailVerified {
		_, _ = w.Write([]byte("Invalid verification link. Perhaps it's expired?"))
		return
	}

	acc.IsEmailVerified = true

	err = IAccountCollection.Update(bson.M{"username": claims["username"]}, acc)
	if err != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte("Internal server error."))
		return
	}

	_, _ = w.Write([]byte("Account successfully verified! Redirecting you to login page...<script>setTimeout(()=>{window.location.replace('" + SITE_DOMAIN + "/auth/login.php')}, 2000)</script>"))
}

 // Get account route
 // GET /v1/accounts/:id
 // https://connectustoday.github.io/api-server/api-reference#accounts

func GetAccountRoute(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
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

	accountapi := interfaces_conv.ConvertToIAccountAPI(account)

	// TODO REDO BY SWITCHING TO USER/ORGANIZATION SPECIFIC OBJ

	b, err := json.Marshal(accountapi)
	if err != nil {
		SendError(w, http.StatusInternalServerError, internalServerError+" (Problem finding account)", 3002)
		return
	}

	_, err = w.Write([]byte(b))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func GetAccountProfileRoute(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var account bson.M
	err := IAccountCollection.Find(bson.M{"username": p.ByName("id")}).One(&account)

	if err != nil {
		if err.Error() == "not found" {
			SendError(w, http.StatusNotFound, notFound+" (Account not found)", 3003)
		} else {
			SendError(w, http.StatusInternalServerError, internalServerError+" (Problem finding account)", 3002)
		}
		return
	}

	acc, err := interfaces_conv.ConvertBSONToIAccount(account)

	if err != nil {
		SendError(w, http.StatusInternalServerError, internalServerError+" (Problem finding account)", 3002)
		return
	}

	var b []byte
	if acc.Type == "User" {
		d, _ := interfaces_conv.ConvertBSONToIUser(account)
		b, err = json.Marshal(d.PersonalInfo)
		if err != nil {
			SendError(w, http.StatusInternalServerError, internalServerError+" (Problem finding account)", 3002)
			return
		}
	} else if acc.Type == "Organization" {
		d, _ := interfaces_conv.ConvertBSONToIOrganization(account)
		b, err = json.Marshal(d.OrgInfo)
		if err != nil {
			SendError(w, http.StatusInternalServerError, internalServerError+" (Problem finding account)", 3002)
			return
		}
	}

	_, err = w.Write([]byte(b))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func GetAccountConnectionsRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func GetAccountPostsRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func RequestConnectionRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func AcceptConnectionRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}