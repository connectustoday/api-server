package main

import (
	"github.com/globalsign/mgo/bson"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"interfaces-conv"
	"interfaces-internal"
	"log"
	"mail-templates"
	"net/http"
	"time"
)

func VerifyUniqueEmail(email string) bool {
	count, err := IAccountCollection.Find(bson.M{"email": email}).Count()
	if err != nil {
		log.Println(err)
		return false
	}
	return count == 0
}

// Account registration route
// POST /v1/auth/register
// https://connectustoday.github.io/api-server/api-reference#register

func RegisterRoute(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	type requestForm struct {
		// Global account fields
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
			log.Println(err.Error())
		}
		SendError(w, http.StatusInternalServerError, internalServerError+" (There was a problem reading the request.)", 3205)
		return
	}
	if !VerifyFieldsExist(&req, FormOmit([]string{"FirstName", "Birthday", "IsNonProfit", "PreferredName"}), false) { // Check request for correct fields
		SendError(w, http.StatusBadRequest, badRequest+" (Bad request.)", 4050)
		return
	}
	if !VerifyUniqueEmail(*req.Email) { // Check if email is unique
		SendError(w, http.StatusBadRequest, badRequest+" (Email already taken.)", 3201)
		return
	}

	// Get hashed bcrypt password
	hashedPassword, err := CreateHashedPassword(*req.Password)
	if err != nil { // check for bcrypt error
		SendError(w, http.StatusInternalServerError, internalServerError+" (There was a problem registering the account.)", 3203)
		return
	}

	verifyToken, err := createEmailVerifyToken(*req.Email)
	if err != nil {
		log.Printf("Problem creating verification token: %s", err.Error())
		SendError(w, http.StatusInternalServerError, internalServerError, 3204)
		return
	}

	// Send email verification email
	if err = sendVerificationEmail(verifyToken, *req.Email); err != nil {
		log.Printf("Problem sending mail: %s", err.Error())
		SendError(w, http.StatusInternalServerError, internalServerError+" (There was a problem sending the verification email. Please ask a website administrator for help.)", 3204)
		return
	}

	if *req.Type == "user" {

		if !VerifyFieldsExist(&req, FormOmit([]string{"FirstName", "Birthday", "IsNonProfit", "PreferredName"}), true) { // Check request for correct fields
			SendError(w, http.StatusBadRequest, badRequest+" (Bad request.)", 4050)
			return
		}

		// Default user
		err = IAccountCollection.Insert(&interfaces_internal.IUser{
			SchemaVersion:        0,
			ID:                   bson.NewObjectId(),
			Email:                *req.Email,
			Password:             string(hashedPassword),
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
			AdminNote:          "",
			Type:               "User",
			PasswordResetToken: "",
			VerifyEmailToken:   verifyToken,
			AuthKey:            GenAuthKey(),
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
			SendError(w, http.StatusBadRequest, badRequest+" (Bad request.)", 4050)
			return
		}

		// Default organization
		err = IAccountCollection.Insert(&interfaces_internal.IOrganization{
			SchemaVersion:        0,
			ID:                   bson.NewObjectId(),
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
			AdminNote:          "",
			Type:               "Organization",
			PasswordResetToken: "",
			VerifyEmailToken:   verifyToken,
			AuthKey:            GenAuthKey(),
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

// create jwt token for account verification

func createEmailVerifyToken(email string) (string, error) {
	// expires in one week
	tok, err := CreateJWTTokenHelper(REGISTER_VERIFY_SECRET, time.Now().Add(time.Second * time.Duration(43200)).Unix(), map[string]interface{}{
		"email": email,
	})
	if err != nil {
		return "", err
	}
	return tok, nil
}

func sendVerificationEmail(token string, email string) error {
	verifyLink := API_DOMAIN + "/v1/auth/verify-email/" + token
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

	err = IAccountCollection.Find(bson.M{"email": claims["email"]}).One(&acc)

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

	if acc.VerifyEmailToken != p.ByName("token") {
		_, _ = w.Write([]byte("Invalid verification link. Perhaps it's expired?"))
		return
	}

	acc.IsEmailVerified = true
	acc.VerifyEmailToken = ""

	err = IAccountCollection.Update(bson.M{"email": claims["email"]}, acc)
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
	account, err := GetAccountDBAsBSON(GetOneAccountQuery(p.ByName("id")))

	if CheckMongoQueryError(w, err, " (Account not found)", 4000, 4001) != nil {
		return
	}

	FetchAccountRouteHelper(account, w)
}

// Get account profile route
// GET /v1/accounts/:id/profile
// https://connectustoday.github.io/api-server/api-reference#accounts

func GetAccountProfileRoute(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	account, err := GetAccountDBAsBSON(GetOneAccountQuery(p.ByName("id")))

	if CheckMongoQueryError(w, err, " (Account not found)", 4000, 4001) != nil {
		return
	}

	FetchProfileRouteHelper(account, w)
}

func GetAccountConnectionsRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func GetAccountPostsRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func RequestConnectionRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params, account bson.M) {

}

func AcceptConnectionRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params, account bson.M) {

}

// Request password reset for user
// POST /v1/accounts/:id/password-reset
// https://connectustoday.github.io/api-server/api-reference#accounts

func RequestPasswordResetRoute(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	type requestForm struct {
		Email *string `json:"email" schema:"email"`
	}

	var req requestForm
	err := DecodeRequest(r, &req)
	if err != nil { // Check decoding error
		SendError(w, http.StatusBadRequest, badRequest+" (Bad request.)", 4050)
		return
	}
	if !VerifyFieldsExist(&req, FormOmit([]string{}), false) { // Check request for correct fields
		SendError(w, http.StatusBadRequest, badRequest+" (Bad request.)", 4050)
		return
	}

	var a bson.M
	err = IAccountCollection.Find(GetOneAccountQuery(*req.Email)).One(&a)
	if CheckMongoQueryError(w, err, " (Account not found.)", 4000, 4001) != nil {
		return
	}

	acc, err := interfaces_conv.ConvertBSONToIAccount(a)

	if err != nil {
		SendError(w, http.StatusInternalServerError, internalServerError, 4001)
		return
	}

	if acc.Email != *req.Email { // email verification failed (not same email)
		WriteOK(w) // fake ok to client TODO DELAY THIS TO FAKE SENDING AN EMAIL SPEED
		return
	}

	// Create jwt token to be used in email
	tok, err := CreateJWTTokenHelper(PASSWORD_RESET_SECRET, time.Now().Add(time.Second * time.Duration(43200)).Unix(), map[string]interface{}{
		"id": acc.ID.Hex(),
	})

	if err != nil {
		SendError(w, http.StatusInternalServerError, internalServerError, 4001)
		return
	}

	// Set password reset token and save to DB

	if acc.Type == "User" {
		user, err := interfaces_conv.ConvertBSONToIUser(a)
		if err != nil {
			SendError(w, http.StatusInternalServerError, internalServerError, 4001)
			return
		}
		user.PasswordResetToken = tok

		err = IAccountCollection.Update(GetOneAccountQuery(*req.Email), user)

	} else if acc.Type == "Organization" {
		org, err := interfaces_conv.ConvertBSONToIOrganization(a)
		if err != nil {
			SendError(w, http.StatusInternalServerError, internalServerError, 4001)
			return
		}
		org.PasswordResetToken = tok

		err = IAccountCollection.Update(GetOneAccountQuery(acc.Email), org)
	}

	if err != nil {
		SendError(w, http.StatusInternalServerError, internalServerError, 4001)
		return
	}

	// Send password reset email

	verifyLink := SITE_DOMAIN + "/password-reset/" + tok
	err = SendMail(acc.Email, "ConnectUS Account Password Reset", mail_templates.FORGOT_PASSWORD, struct {
		Site template.URL
		Account string
		ResetLink template.URL
	}{Site: template.URL(SITE_DOMAIN), Account: acc.Email, ResetLink: template.URL(verifyLink)})

	if err != nil {
		log.Printf("Problem sending mail: %s", err.Error())
		SendError(w, http.StatusInternalServerError, internalServerError, 4001)
		return
	}

	WriteOK(w)
}
