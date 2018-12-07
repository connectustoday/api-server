/*
 *
 *     Copyright (C) 2018 ConnectUS
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU General Public License for more details.
 *
 *     You should have received a copy of the GNU General Public License
 *     along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package main

import (
	"encoding/json"
	"github.com/gorilla/schema"
	"net/http"
	"reflect"
	"strings"
)

var (
	schemaDecoder = schema.NewDecoder()
)

func init() {
	schemaDecoder.IgnoreUnknownKeys(true)
}

func FormOmit(omitFields []string) (ret map[string]bool) {
	ret = make(map[string]bool)
	for _, e := range omitFields {
		ret[e] = true
	}
	return
}

func VerifyFieldsExist(obj interface{}, omitFields map[string]bool) bool {
	v := reflect.ValueOf(obj)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).IsNil() && !omitFields[v.Type().Field(i).Name] {
			return false
		}
	}
	return true
}

func WriteOK(w http.ResponseWriter) (err error) {
	_, err = w.Write([]byte(`{"message": "` + ok + `"}`))
	return
}

func DecodeRequest(r *http.Request, obj interface{}) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}

	if strings.EqualFold(r.Header.Get("Content-Type"), "application/x-www-form-urlencoded") || strings.EqualFold(r.Header.Get("Content-Type"), "application/form-data") {
		return schemaDecoder.Decode(obj, r.PostForm)
	} else {
		return json.NewDecoder(r.Body).Decode(&obj)
	}
}