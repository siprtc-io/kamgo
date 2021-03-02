package api

import (
	"fmt"
	"gopkg.in/go-playground/validator.v9"
	"kamgo/model"
	"kamgo/modules/log"
	"net/http"
	"strconv"
	"strings"
)

//Create New SipEndpoint
func CreateSipEndpoint(c *Context) error {
	var sipEndPoint model.KamSubscriber
	req := &sipEndpointCreateRequest{}
	var errTemplate model.ErrorTemplate
	errTemplate.Host = model.HostSanity(c.Request().Host)
	errTemplate.Resource = c.Request().RequestURI

	if err := req.bind(c, &sipEndPoint); err != nil {
		log.Error(err)
		for _, err := range err.(validator.ValidationErrors) {
			errTemplate.Resource = err.StructField()
			break
		}
		return c.JSON(http.StatusBadRequest, model.GetErrorStruct("INVALID-RESOURCE", errTemplate))
	}

	if err := sipEndPoint.Create(); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.GetErrorStruct("RESOURCE-NOT-FOUND", errTemplate))
	} else {
		return c.JSON(http.StatusCreated, sipEndPoint)
	}
}

//Delete sip endpoint
func DeleteSipEndpoint(c *Context) error {
	endpoint := c.Param("endpoint")
	var errTemplate model.ErrorTemplate
	errTemplate.Host = model.HostSanity(c.Request().Host)
	errTemplate.Resource = c.Request().RequestURI
	var sipEndPoint model.KamSubscriber

	if endpoint == "" {
		return c.JSON(http.StatusBadRequest, model.GetErrorStruct("INVALID-RESOURCE", errTemplate))
	}

	endpointSplit := strings.FieldsFunc(endpoint, Split)

	sipPrefix := ""
	sipDomain := ""
	sipUser := ""

	for i := range endpointSplit {
		switch i {
		case 0:
			sipPrefix = strings.ToLower(endpointSplit[0])
			if sipPrefix != "sip" {
				return c.JSON(http.StatusBadRequest, model.GetErrorStruct("INVALID-RESOURCE", errTemplate))
			}
		case 1:
			sipUser = strings.ToLower(endpointSplit[1])
		case 2:
			sipDomain = strings.ToLower(endpointSplit[2])
		default:

		}
	}

	if sipUser == "" || sipDomain == "" {
		return c.JSON(http.StatusBadRequest, model.GetErrorStruct("INVALID-RESOURCE", errTemplate))
	}

	if err := sipEndPoint.DeleteKamSubModel(sipDomain, sipUser); err != nil {
		return c.JSON(http.StatusBadRequest, model.GetErrorStruct("INVALID-RESOURCE", errTemplate))
	}
	return c.JSON(http.StatusNoContent, "")
}

func GetSipEndpoint(c *Context) error {
	endpoint := c.Param("endpoint")
	var errTemplate model.ErrorTemplate
	errTemplate.Host = model.HostSanity(c.Request().Host)
	errTemplate.Resource = c.Request().RequestURI
	var sipEndPoint model.KamSubscriber

	if endpoint == "" {
		return c.JSON(http.StatusBadRequest, model.GetErrorStruct("INVALID-RESOURCE", errTemplate))
	}

	endpointSplit := strings.FieldsFunc(endpoint, Split)

	sipPrefix := ""
	sipDomain := ""
	sipUser := ""

	for i := range endpointSplit {
		switch i {
		case 0:
			sipPrefix = strings.ToLower(endpointSplit[0])
			if sipPrefix != "sip" {
				return c.JSON(http.StatusBadRequest, model.GetErrorStruct("INVALID-RESOURCE", errTemplate))
			}
		case 1:
			sipUser = strings.ToLower(endpointSplit[1])
		case 2:
			sipDomain = strings.ToLower(endpointSplit[2])
		default:

		}
	}

	if sipUser == "" || sipDomain == "" {
		return c.JSON(http.StatusBadRequest, model.GetErrorStruct("INVALID-RESOURCE", errTemplate))
	}

	if err := sipEndPoint.GetKamSubModel(sipDomain, sipUser); err != nil {
		return c.JSON(http.StatusBadRequest, model.GetErrorStruct("INVALID-RESOURCE", errTemplate))
	}
	return c.JSON(http.StatusOK, sipEndPoint)
}

func GetListSipEndpoint(c *Context) error {
	var errTemplate model.ErrorTemplate
	errTemplate.Host = model.HostSanity(c.Request().Host)
	errTemplate.Resource = c.Request().RequestURI
	var sipEndPoint model.KamSubscriber

	pageNumber, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || pageNumber == 0 {
		pageNumber = 1
	}
	requestURI := fmt.Sprintf("%s%s", c.Request().Host, c.Request().RequestURI)
	sipEndpointList := sipEndPoint.GetAllSipEndpoints(requestURI, pageNumber)
	return c.JSON(http.StatusOK, sipEndpointList)
}

func UpdateSipEndpoint(c *Context) error {
	endpoint := c.Param("endpoint")
	var errTemplate model.ErrorTemplate
	errTemplate.Host = model.HostSanity(c.Request().Host)
	errTemplate.Resource = c.Request().RequestURI
	var sipEndPoint model.KamSubscriber

	if endpoint == "" {
		return c.JSON(http.StatusBadRequest, model.GetErrorStruct("INVALID-RESOURCE", errTemplate))
	}

	endpointSplit := strings.FieldsFunc(endpoint, Split)

	sipPrefix := ""
	sipDomain := ""
	sipUser := ""

	for i := range endpointSplit {
		switch i {
		case 0:
			sipPrefix = strings.ToLower(endpointSplit[0])
			if sipPrefix != "sip" {
				return c.JSON(http.StatusBadRequest, model.GetErrorStruct("INVALID-RESOURCE", errTemplate))
			}
		case 1:
			sipUser = strings.ToLower(endpointSplit[1])
		case 2:
			sipDomain = strings.ToLower(endpointSplit[2])
		default:

		}
	}
	if sipUser == "" || sipDomain == "" {
		return c.JSON(http.StatusBadRequest, model.GetErrorStruct("INVALID-RESOURCE", errTemplate))
	}
	req := &sipEndpointUpdateRequest{}
	if err := req.bind(c, &sipEndPoint); err != nil {
		log.Error(err)
		for _, err := range err.(validator.ValidationErrors) {
			errTemplate.Resource = err.StructField()
			break
		}
		return c.JSON(http.StatusBadRequest, model.GetErrorStruct("INVALID-RESOURCE", errTemplate))
	}
	if err := sipEndPoint.KamUpdate(sipDomain, sipUser); err != nil {
		return c.JSON(http.StatusBadRequest, model.GetErrorStruct("INVALID-RESOURCE", errTemplate))
	}
	return c.JSON(http.StatusAccepted, sipEndPoint)
}
func Split(r rune) bool {
	return r == ':' || r == '@' || r == ';'
}
