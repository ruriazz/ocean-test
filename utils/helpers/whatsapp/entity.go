package whatsappHelper

type WhatsAppTemplateMessage struct {
	MessagingProduct string           `json:"messaging_product"`
	To               string           `json:"to"`
	Type             string           `json:"type"`
	Template         WhatsAppTemplate `json:"template"`
}

type WhatsAppTemplate struct {
	Name       string                      `json:"name"`
	Language   WhatsAppTemplateLanguage    `json:"language"`
	Components []WhatsAppTemplateComponent `json:"components"`
}

type WhatsAppTemplateLanguage struct {
	Code string `json:"code"`
}

type WhatsAppTemplateComponent struct {
	Type       string                      `json:"type"`
	Parameters []WhatsAppTemplateParameter `json:"parameters"`
}

type WhatsAppTemplateParameter struct {
	Type string `json:"type"`
	Text string `json:"text"`
}
