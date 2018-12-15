package main

import (
	"bytes"
	"encoding/json"
	"github.com/globalsign/mgo/bson"
	"github.com/julienschmidt/httprouter"
	"interfaces-api"
	"interfaces-conv"
	"interfaces-internal"
	"io/ioutil"
	"net/http"
	"reflect"
)

// Get personal account profile route
// GET /v1/profile
// https://connectustoday.github.io/api-server/api-reference#accounts

func GetPersonalAccountProfileRoute(w http.ResponseWriter, _ *http.Request, _ httprouter.Params, account bson.M) {
	w.Header().Set("Content-Type", "application/json")
	FetchProfileRouteHelper(account, w)
}

// Get personal account route
// GET /v1/account
// https://connectustoday.github.io/api-server/api-reference#accounts

func GetPersonalAccountRoute(w http.ResponseWriter, _ *http.Request, _ httprouter.Params, account bson.M) {
	w.Header().Set("Content-Type", "application/json")
	FetchAccountRouteHelper(account, w)
}

// Get personal account settings route
// GET /v1/settings
// https://connectustoday.github.io/api-server/api-reference#accounts

func GetPersonalAccountSettingsRoute(w http.ResponseWriter, _ *http.Request, _ httprouter.Params, account bson.M) {
	w.Header().Set("Content-Type", "application/json")
	acc, err := interfaces_conv.ConvertBSONToIAccount(account)

	if err != nil {
		SendError(w, http.StatusInternalServerError, internalServerError+" (Problem finding account)", 4001)
		return
	}

	var b []byte
	if acc.Type == "User" {
		d, _ := interfaces_conv.ConvertBSONToIUser(account)
		b, err = json.Marshal(interfaces_conv.ConvertToIUserSettingsAPI(d.Settings, d.Type))
	} else if acc.Type == "Organization" {
		d, _ := interfaces_conv.ConvertBSONToIOrganization(account)
		b, err = json.Marshal(interfaces_conv.ConvertToIOrganizationSettingsAPI(d.Settings, d.Type))
	}
	if err != nil {
		SendError(w, http.StatusInternalServerError, internalServerError+" (Problem finding account)", 4001)
		return
	}

	_, err = w.Write([]byte(b))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// Change account settings route
// PATCH /v1/settings
// https://connectustoday.github.io/api-server/api-reference#accounts

func PatchAccountSettingsRoute(w http.ResponseWriter, r *http.Request, _ httprouter.Params, account bson.M) {
	w.Header().Set("Content-Type", "application/json")

	acc, err := interfaces_conv.ConvertBSONToIAccount(account)
	if err != nil {
		SendError(w, http.StatusInternalServerError, internalServerError, 4001)
		return
	}

	// save original body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		SendError(w, http.StatusInternalServerError, internalServerError, 4001)
		return
	}
	// restore body for first reading
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	var save interface{}

	if acc.Type == "User" {

		var user interfaces_internal.IUser
		user, err := interfaces_conv.ConvertBSONToIUser(account)
		if err != nil {
			SendError(w, http.StatusInternalServerError, internalServerError, 4001)
			return
		}

		var profile interfaces_api.IUserSettingsAPI
		// parse into api object (cast field types and type checks)
		err = DecodeRequest(r, &profile)
		if err != nil {
			SendError(w, http.StatusBadRequest, internalServerError + " when parsing request. Bad input?", 4050)
			return
		}
		// restore body for second reading
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		// Patch object
		err = PatchObjectUsingRequest(w, r, reflect.ValueOf(&profile).Elem(), reflect.ValueOf(&user.Settings).Elem())
		if err != nil {
			SendError(w, http.StatusInternalServerError, internalServerError, 4001)
			return
		}
		save = user // save patched object

	} else if acc.Type == "Organization" {

		var org interfaces_internal.IOrganization
		org, err := interfaces_conv.ConvertBSONToIOrganization(account)
		if err != nil {
			SendError(w, http.StatusInternalServerError, internalServerError, 4001)
			return
		}

		var profile interfaces_api.IOrganizationSettingsAPI
		// parse into api object (cast field types and type checks)
		err = DecodeRequest(r, &profile)
		if err != nil {
			SendError(w, http.StatusBadRequest, internalServerError + " when parsing request. Bad input?", 4050)
			return
		}
		// restore body for second reading
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		// Patch object
		err = PatchObjectUsingRequest(w, r, reflect.ValueOf(&profile).Elem(), reflect.ValueOf(&org.Settings).Elem())
		if err != nil {
			SendError(w, http.StatusInternalServerError, internalServerError, 4001)
			return
		}
		save = org // save patched object
	}

	// Save updated profile to DB
	err = IAccountCollection.Update(bson.M{"username": acc.UserName}, save)
	if err != nil {
		SendError(w, http.StatusInternalServerError, internalServerError, 4001)
		return
	}
	WriteOK(w)
}

// Change account profile route
// PATCH /v1/profile
// https://connectustoday.github.io/api-server/api-reference#accounts

func PatchAccountProfileRoute(w http.ResponseWriter, r *http.Request, _ httprouter.Params, account bson.M) {
	w.Header().Set("Content-Type", "application/json")

	acc, err := interfaces_conv.ConvertBSONToIAccount(account)
	if err != nil {
		SendError(w, http.StatusInternalServerError, internalServerError, 4001)
		return
	}

	// save original body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		SendError(w, http.StatusInternalServerError, internalServerError, 4001)
		return
	}
	// restore body for first reading
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	var save interface{}

	if acc.Type == "User" {

		var user interfaces_internal.IUser
		user, err := interfaces_conv.ConvertBSONToIUser(account)
		if err != nil {
			SendError(w, http.StatusInternalServerError, internalServerError, 4001)
			return
		}

		var profile interfaces_api.IUserProfileAPI
		// parse into api object (cast field types and type checks)
		err = DecodeRequest(r, &profile)
		if err != nil {
			SendError(w, http.StatusBadRequest, internalServerError + " when parsing request. Bad input?", 4050)
			return
		}
		// restore body for second reading
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		// Patch object
		err = PatchObjectUsingRequest(w, r, reflect.ValueOf(&profile).Elem(), reflect.ValueOf(&user.PersonalInfo).Elem())
		if err != nil {
			SendError(w, http.StatusInternalServerError, internalServerError, 4001)
			return
		}
		save = user // save patched object

	} else if acc.Type == "Organization" {

		var org interfaces_internal.IOrganization
		org, err := interfaces_conv.ConvertBSONToIOrganization(account)
		if err != nil {
			SendError(w, http.StatusInternalServerError, internalServerError, 4001)
			return
		}

		var profile interfaces_api.IOrganizationProfileAPI
		// parse into api object (cast field types and type checks)
		err = DecodeRequest(r, &profile)
		if err != nil {
			SendError(w, http.StatusBadRequest, internalServerError + " when parsing request. Bad input?", 4050)
			return
		}
		// restore body for second reading
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		// Patch object
		err = PatchObjectUsingRequest(w, r, reflect.ValueOf(&profile).Elem(), reflect.ValueOf(&org.OrgInfo).Elem())
		if err != nil {
			SendError(w, http.StatusInternalServerError, internalServerError, 4001)
			return
		}
		save = org // save patched object
	}

	// Save updated profile to DB
	err = IAccountCollection.Update(bson.M{"username": acc.UserName}, save)
	if err != nil {
		SendError(w, http.StatusInternalServerError, internalServerError, 4001)
		return
	}
	WriteOK(w)
}