package incomsms

import (
	"fmt"
	"log"
	"testing"
)

func setup(t *testing.T) *Client {
	username := "username"
	password := "password"

	client, err := NewClient(username, password)
	if err != nil {
		t.Fatal(err)
	}
	client.logger = log.Println
	return client
}

func TestSendMessage(t *testing.T) {
	client := setup(t)

	brandName := "Brand"
	result, _, err := client.Sms.SendMessage(SendMessageInput{
		PhoneNumber:      "84932123456",
		MsgContent:       fmt.Sprintf("%s gui ma OTP dang nhap ung dung Dich Vu %s cua ban la 295788. Ma xac minh co hieu luc trong 3 phut.", brandName, brandName),
		PrefixID:         brandName,
		CommandCode:      brandName,
		RequestID:        "0",
		FeeTypeID:        0,
		MsgContentTypeID: 0,
	})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result.StatusCode != CodeInvalidUsernameOrPasswordOrIP {
		t.Fatalf("CodeResult: expected %s, got %s", CodeSuccess, result.StatusCode)
	}
}
