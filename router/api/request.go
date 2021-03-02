package api

import (
	"kamgo/model"
)

type sipEndpointCreateRequest struct {
	Username string `json:"username" validate:"required,alphanum,gt=1"`
	Password string `json:"password" validate:"required,gt=5"`
	Domain   string `json:"domain" validate:"required,fqdn"`
}

type sipEndpointUpdateRequest struct {
	Username string `json:"username" validate:"required,alphanum,gt=1"`
	Password string `json:"password" validate:"required,gt=5"`
}

func (sipReqEndpoint *sipEndpointCreateRequest) bind(c *Context, sip *model.KamSubscriber) error {
	if err := c.Bind(sipReqEndpoint); err != nil {
		return err
	}
	if err := c.Validate(sipReqEndpoint); err != nil {
		return err
	}
	sip.Username = sipReqEndpoint.Username
	sip.Password = sipReqEndpoint.Password
	sip.DomainName = sipReqEndpoint.Domain
	return nil
}

func (sipReqEndpoint *sipEndpointUpdateRequest) bind(c *Context, sip *model.KamSubscriber) error {
	if err := c.Bind(sipReqEndpoint); err != nil {
		return err
	}
	if err := c.Validate(sipReqEndpoint); err != nil {
		return err
	}
	sip.Username = sipReqEndpoint.Username
	sip.Password = sipReqEndpoint.Password
	return nil
}
