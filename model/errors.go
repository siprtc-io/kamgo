package model

import (
	"bytes"
	"strings"
	"text/template"
)

type ErrorSipServiceResponse struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	MoreInfo string `json:"more_info"`
	Status   int    `json:"status"`
}

type ErrorTemplate struct {
	Host   string
	Resource string
}

var SipServiceValidation = map[string]ErrorSipServiceResponse{
	"INVALID-RESOURCE": {
		Code:     21213,
		Message:  "{{.Resource}} is not valid",
		MoreInfo: "https://{{.Host}}",
		Status:   400,
	},
	"MALFORMED": {
		Code:     21210,
		Message:  "Request is not able to decode, Please check and send again",
		MoreInfo: "https://{{.Host}}",
		Status:   400,
	},
	"RESOURCE-NOT-FOUND": {
		Code:     20404,
		Message:  "The requested resource {{.Resource}} was not found",
		MoreInfo: "https://{{.Host}}",
		Status:   404,
	},
	"BODY-PARAM-INCORRECT": {
		Code:     20001,
		Message:  "{{.Resource}} is not a valid choice",
		MoreInfo: "https://{{.Host}}",
		Status:   400,
	},
}

func GetErrorStruct(state string, errorTemplate interface{}) ErrorSipServiceResponse {
	validationError := SipServiceValidation[state]
	validationError.Message = parseTemplate(validationError.Message, errorTemplate)
	validationError.MoreInfo = parseTemplate(validationError.MoreInfo, errorTemplate)
	return validationError
}

func parseTemplate(tmpl string, errorTemplate interface{}) string {
	var err error
	// with name passed as argument
	sipValidationTmpl := template.New("SipValidation")
	// "Parse" parses a string into a template
	sipValidationTmpl, err = sipValidationTmpl.Parse(tmpl)
	if err != nil {
		return tmpl
	}
	var tmplBytes bytes.Buffer

	if err = sipValidationTmpl.Execute(&tmplBytes, errorTemplate); err != nil {
		return tmpl
	}
	return tmplBytes.String()
}

func HostSanity(host string) string {
	return strings.Replace(host, "api.", "", -1)
}
