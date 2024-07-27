package domainEntity

type CreateOtpData struct {
	WhatsappNumber string `json:"whatsapp_number" validate:"required"`
	Token          string `json:"token"`
}

type CheckOtpData struct {
	WhatsappNumber string `json:"whatsapp_number" validate:"required"`
	Token          string `json:"token" validate:"required"`
	Otp            string `json:"otp" validate:"required"`
	Valid          bool   `json:"valid"`
}
