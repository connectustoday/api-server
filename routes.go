package api_server

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func ExperienceRoutes(prefix string, router *httprouter.Router) {

}

func OpportunityRoutes(prefix string, router *httprouter.Router) {

}

func AuthRoutes(prefix string, router *httprouter.Router) {

	router.GET(prefix, func (w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

         router.POST(prefix + "/register", WithAccountVerify(RegisterRoute))

         /*
          * Test utility to check if logged in
          */

          if DEBUG {
          		router.GET(prefix + "/me", WithAccountVerify(func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
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
}

func AccountRoutes(prefix string, router *httprouter.Router) {
	router.GET(prefix, func (w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		//TODO
	})

	/*
	 * Global Account Routes
	 */

	router.GET(prefix + "/search", func (w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	})

	// Fetch an Account's basic information (Global)

	router.GET(prefix + "/:id")
}

func PersonalAccountsRoutes(prefix string, router *httprouter.Router) {

}