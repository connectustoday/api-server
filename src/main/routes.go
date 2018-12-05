package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"interfaces-internal"
	"net/http"
)

type accountPassRoute func(w http.ResponseWriter, r *http.Request, _ httprouter.Params, account interfaces_internal.IAccount)

/*
 * Routes for experiences
 */

func ExperienceRoutes(prefix string, router *httprouter.Router) {
	/*
	 * Experience Routes
	 */

	// Get current user's experiences
	// No parameters required (except header token)
	router.GET(prefix, WithAccountVerify(GetPersonalExperiencesRoute))

	// Create experience
	// Use IExperienceAPI object as "experience" field
	router.POST(prefix, WithAccountVerify(CreateExperienceRoute))

	// Update experience
	// Use IExperienceAPI object as "experience" field for the new replacing experience
	router.PUT(prefix+"/:id", WithAccountVerify(UpdateExperienceRoute))

	// Delete experience
	// No parameters required (except header token)
	router.DELETE(prefix+"/:id", WithAccountVerify(DeleteExperienceRoute))

	// List pending experience validations (for organization)
	// No parameters required (except header token)
	router.GET(prefix+"/validations", WithAccountVerify(GetExperienceValidationsRoute))

	// Approve or don't approve validation (for organization)
	// approve = boolean on whether or not to approved the validation
	router.POST(prefix+"/validations/:user/:id", WithAccountVerify(ReviewExperienceValidationsRoute))

	// Approve validation (from email)
	router.GET(prefix+"/email_approve/:token", WithAccountVerify(EmailApproveExperienceValidationRoute))
}

/*
 * Routes for opportunities
 */

func OpportunityRoutes(prefix string, router *httprouter.Router) {

}

/*
 * Routes for authentication and account creation
 */

func AuthRoutes(prefix string, router *httprouter.Router) {

	router.GET(prefix, func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		_, err := w.Write([]byte(badRequest))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

	// Register Account API Endpoint
	// https://connectustoday.github.io/api-server/api-reference#register

	router.POST(prefix+"/register", RegisterRoute)

	/*
	 * Login API Endpoint
	 * https://connectustoday.github.io/api-server/api-reference#login
     * TODO EMAIL LOGIN (RATHER THAN USERNAME)
     */

	router.POST(prefix+"/login", LoginRoute)

	// Verify email using jsonwebtoken

	router.GET(prefix+"/verify-email/:token", VerifyEmailRequestRoute)

	/*
	 * Test utility to check if logged in
	 */

	if DEBUG {
		router.GET(prefix+"/me", WithAccountVerify(func(writer http.ResponseWriter, request *http.Request, params httprouter.Params, account interfaces_internal.IAccount) {
			vjson, err := json.Marshal(account)
			if err != nil {
				println(err)
			}
			_, err2 := writer.Write(vjson)
			if err2 != nil {
				println(err)
				writer.WriteHeader(500)
			}
		}))
	}
}

func AccountRoutes(prefix string, router *httprouter.Router) {
	router.GET(prefix, func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		_, _ = w.Write([]byte(badRequest))
	})

	/*
	 * Global Account Routes
	 */

	router.GET(prefix+"/search", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	})

	// Fetch an Account's basic information (Global)

	router.GET(prefix+"/:id", GetAccountRoute)

	// Get the profile of a user

	router.GET(prefix + "/:id/profile", GetAccountProfileRoute)

	// Get the connections of a user

	router.GET(prefix + "/:id/connections", GetAccountConnectionsRoute)

	// Get the experiences list of a user

	router.GET(prefix + "/:id/experiences", GetExperiencesRoute)

}

func PersonalAccountsRoutes(prefix string, router *httprouter.Router) {

}
