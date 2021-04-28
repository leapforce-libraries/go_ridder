package ridder

import (
	"encoding/base64"
	"encoding/xml"
	"strconv"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type inboundXMLMessage struct {
	MessageID              string `json:"MessageId"`
	Base64EncodedXMLString string `json:"Base64EncodedXmlString"`
}

func (service *Service) SendXMLMessage(messageID string, object interface{}) (*int32, *errortools.Error) {
	b, err := xml.Marshal(object)
	if err != nil {
		return nil, errortools.ErrorMessage(err)
	}
	base64encodedXMLString := base64.StdEncoding.EncodeToString(b)

	message := inboundXMLMessage{
		MessageID:              messageID,
		Base64EncodedXMLString: base64encodedXMLString,
	}

	var idString string
	requestConfig := go_http.RequestConfig{
		URL:           service.url("inboundxmlmessage"),
		BodyModel:     message,
		ResponseModel: &idString,
	}
	_, _, e := service.post(&requestConfig)
	if e != nil {
		return nil, e
	}

	var id *int32 = nil
	if idString != "" {
		idInt64, err := strconv.ParseInt(idString, 10, 32)
		if err != nil {
			return nil, errortools.ErrorMessage(err)
		}

		idInt32 := int32(idInt64)
		id = &idInt32
	}

	return id, nil
}
