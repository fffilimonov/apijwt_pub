package services

import (
	"../services/models"
	b64 "encoding/base64"
	. "github.com/mailjet/mailjet-apiv3-go"
	"net/http"
	"fmt"
)

func SendActivation(requestUser *models.User) (int, []byte) {

	sEnc := b64.URLEncoding.EncodeToString([]byte(requestUser.Username))

	mailjetClient := NewMailjetClient("", "")
	email := &InfoSendMail{
		FromEmail: "admin@likeuser.com",
		FromName:  "Activate Mail",
		Subject:   "Activate Likeuser.com account",
		TextPart:  "Hello.\nLink for activation:\n" + "https://likeuser.com/activate/" + sEnc + "\nThank you.",
		Recipients: []Recipient{
			Recipient{
				Email: requestUser.Username,
			},
		},
	}
	_, err := mailjetClient.SendMail(email)
	fmt.Printf("link: %v\n", "https://likeuser.com/activate/" + sEnc)
	if err != nil {
		panic(err)
	} else {
		return http.StatusOK, []byte("")
	}
}
