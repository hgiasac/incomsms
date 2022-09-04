package incomsms

import (
	"net/http"
	"net/url"
)

// SendMessageInput represents the message request input
type SendMessageInput struct {
	Credential
	// message content
	MsgContent string `json:"MsgContent"`
	// phone number of the recipient
	PhoneNumber string `json:"PhoneNumber"`
	// Service number (6x89, 996, 997, 998, 19001255) or the brand name of partner
	PrefixID string `json:"PrefixId,omitempty"`
	// Service code is used for report and management. Can be the brand name
	CommandCode string `json:"CommandCode"`
	// The auto-generated request id. Should be 0 if the message is sent from MT
	RequestID string `json:"RequestId,omitempty"`
	// Set 0 if the message is sent with brand name
	FeeTypeID FeeType `json:"FeeTypeId"`
	// send message content with or without unicode
	MsgContentTypeID MsgContentType `json:"MsgContentTypeId"`
}

// SendMessageResponse represents the send message response
type SendMessageResponse struct {
	StatusCode StatusCode `json:"StatusCode"`
	StatusDesc string     `json:"StatusDesc,omitempty"`
}

type smsService struct {
	client *httpClient
}

// SendMessage send SMS message to the target phone number
func (ss *smsService) SendMessage(input SendMessageInput) (*SendMessageResponse, *http.Response, error) {
	u, err := url.Parse("/MtService/SendSms")
	if err != nil {
		return nil, nil, err
	}
	// create the request
	input.Credential = ss.client.credential
	req, err := ss.client.NewRequest("POST", u.String(), input)
	if err != nil {
		return nil, nil, err
	}

	result := &SendMessageResponse{}
	resp, err := ss.client.Do(req, result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, err
}
