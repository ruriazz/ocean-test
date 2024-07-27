package whatsappHelper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type WhatsappHelper interface {
	SendOtp(recipient string, otp string, validity string) error
}

type helper struct {
	apiUrl      string
	accessToken string
}

func New(businessId string, accessToken string) WhatsappHelper {
	return helper{
		apiUrl:      fmt.Sprintf("https://graph.facebook.com/v20.0/%s/messages", businessId),
		accessToken: fmt.Sprintf("Bearer %s", accessToken),
	}
}

func (h helper) SendOtp(recipient string, otp string, validity string) error {
	otpCode := fmt.Sprintf("```%s```", otp)

	message := WhatsAppTemplateMessage{
		MessagingProduct: "whatsapp",
		To:               recipient,
		Type:             "template",
		Template: WhatsAppTemplate{
			Name: "debug_template",
			Language: WhatsAppTemplateLanguage{
				Code: "id",
			},
			Components: []WhatsAppTemplateComponent{
				{
					Type: "body",
					Parameters: []WhatsAppTemplateParameter{
						{
							Type: "text",
							Text: otpCode,
						},
						{
							Type: "text",
							Text: validity,
						},
					},
				},
			},
		},
	}

	messageJSON, err := json.Marshal(message)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", h.apiUrl, bytes.NewBuffer(messageJSON))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", h.accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
