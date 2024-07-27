package domainInterface

import (
	"time"

	"github.com/gofiber/fiber/v2"
	domainEntity "github.com/ruriazz/ocean-test/api/otp/domain/entity"
)

type OtpHandler interface {
	CreateOtp(ctx *fiber.Ctx) error
	CheckOtp(ctx *fiber.Ctx) error
}

type OtpUsecase interface {
	CreateOtp(requestid string, data *domainEntity.CreateOtpData) error
	CheckOtp(data *domainEntity.CheckOtpData) error
}

type OtpRepository interface {
	SaveOtp(requestid string, phoneNumber string, value string) error
	NextCreateOtpTime(phoneNumber string) (time.Duration, error)
	OtpErrorCounter(phoneNumber string) error
	MaxCheckOtpAttempt(phoneNumber string) bool
	GetOtp(phoneNumber string) (string, time.Time, error)
	UnsetOtp(phoneNumber string)
}
