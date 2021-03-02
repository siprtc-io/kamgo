package api


type sipEndpointResponse struct {
	SipEndpoint struct {
		Name string `json:"name" validate:"required"`
		Username string `json:"username" validate:"required"`
		DomainName string `json:"username" validate:"required"`
		CodecList      string  `json:"codec_list" validate:"required"`
		IsActive bool `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	} `json:"sipEndpoint"`
}

