package api_server

import (
	"github.com/julienschmidt/httprouter"
	"interfaces_internal"
	"net/http"
)

func verifyUniqueUsername(username string) {
	var results []interfaces_internal.IAccount
}

func getAccount(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	
}