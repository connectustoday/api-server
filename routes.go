package api_server

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func ExperienceRoutes(prefix string, router *httprouter.Router) {

}

func OpportunityRoutes(prefix string, router *httprouter.Router) {

}

func AuthRoutes(prefix string, router *httprouter.Router) {

}

func AccountRoutes(prefix string, router *httprouter.Router) {
	router.GET(prefix, func (w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprintf(w, "hi", "HI")
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