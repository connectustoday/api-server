package api_server

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	ok                   = "200 - OK"
	created              = "201 - Created"
	accepted             = "202 - Accepted"
	badRequest           = "400 - Bad Request"
	unauthorized         = "401 - Unauthorized"
	forbidden            = "403 - Forbidden"
	notFound             = "404 - Not Found"
	methodNotAllowed     = "405 - Method Not Allowed"
	unacceptable         = "406 - Unacceptable"
	payloadTooLarge      = "413 - Payload Too Large"
	unsupportedMediaType = "415 - Unsupported Media Type"
	tooManyRequests      = "429 - Too Many Requests"
	internalServerError  = "500 - Internal Server Error"
	notImplemented       = "501 - Not Implemented"
	serviceUnavailable   = "503 - Service Unavailable"
)

type ReturnError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func SendError(w http.ResponseWriter, httpCode int, message string, errorCode int) {
	res, err := json.MarshalIndent(ReturnError{Code: errorCode, Message: message}, "", "  ")
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(httpCode)
	_, err = w.Write(res)
	if err != nil {
		log.Println(err)
	}
}
