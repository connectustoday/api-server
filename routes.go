package api_server

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

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
 * Routes for authentication and account creationg
 */

func AuthRoutes(prefix string, router *httprouter.Router) {

	router.GET(prefix, func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		_, err := w.Write([]byte(badRequest))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

	/*
         * Register Endpoint Required Fields
         * - username
         * - email
         * - password
         * - type ("organization", "user")
         * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
         * User Required Fields
         * - first_name
         * - birthday
         * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
         * Organization Required Fields
         * - is_nonprofit
         * - preferred_name
         * ------------------------------------
         * Returns 200 if SUCCESSFUL
         */

	router.POST(prefix+"/register", WithAccountVerify(RegisterRoute))

	/*
	 * Test utility to check if logged in
	 */

	if DEBUG {
		router.GET(prefix+"/me", WithAccountVerify(func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
			vjson, err := json.Marshal(GetAccountFromContext(request.Context()))
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

	/*
	 * Login Endpoint Required Fields
	 * - username
  	 * - password
	 * Returns 200 + token if SUCCESSFUL
     * TODO EMAIL LOGIN (RATHER THAN USERNAME)
     */

	router.POST(prefix+"/login", LoginRoute)

	// Verify email using jsonwebtoken

	router.GET(prefix+"/verify-email/:token", VerifyEmailRequestRoute)
}

func AccountRoutes(prefix string, router *httprouter.Router) {
	router.GET(prefix, func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		//TODO
	})

	/*
	 * Global Account Routes
	 */

	router.GET(prefix+"/search", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	})

	// Fetch an Account's basic information (Global)

	router.GET(prefix+"/:id", GetAccountRoute)

	
}

func PersonalAccountsRoutes(prefix string, router *httprouter.Router) {

}
