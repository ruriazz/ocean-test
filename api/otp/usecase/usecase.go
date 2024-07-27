package usecase

import (
	"fmt"
	"strconv"

	domainConst "github.com/ruriazz/ocean-test/api/otp/domain/const"
	domainEntity "github.com/ruriazz/ocean-test/api/otp/domain/entity"
	domainInterface "github.com/ruriazz/ocean-test/api/otp/domain/interface"
	"github.com/ruriazz/ocean-test/api/otp/repository"
	"github.com/ruriazz/ocean-test/pkg/manager"
	passwordHelper "github.com/ruriazz/ocean-test/utils/helpers/password"
	stringHelper "github.com/ruriazz/ocean-test/utils/helpers/string"
	whatsappHelper "github.com/ruriazz/ocean-test/utils/helpers/whatsapp"
	"github.com/sirupsen/logrus"
)

type usecase struct {
	manager    manager.Manager
	repository domainInterface.OtpRepository
}

func New(manager manager.Manager) domainInterface.OtpUsecase {
	return usecase{
		manager:    manager,
		repository: repository.New(manager),
	}
}

func (u usecase) CreateOtp(requestid string, data *domainEntity.CreateOtpData) error {
	phoneNumber, err := stringHelper.FormatIDNPhoneNumber(data.WhatsappNumber)
	if err != nil {
		return err
	}

	dur, err := u.repository.NextCreateOtpTime(phoneNumber)
	if err != nil {
		u.manager.Log().Error(err)
		return nil
	}

	if dur > 0 {
		return fmt.Errorf("OTP request limit has been exceeded. Please wait %s before trying again", stringHelper.DurationToMinuteString(dur))
	}

	otp := stringHelper.GenerateOTP()
	token := stringHelper.GenerateToken(40)
	hash := u.otpHashing(otp, token)

	if err := u.repository.SaveOtp(requestid, phoneNumber, hash); err != nil {
		u.manager.Log().Error(err)
		return nil
	}

	u.sendOtpToWhatsapp(phoneNumber, otp)

	data.WhatsappNumber = phoneNumber
	data.Token = token
	return nil
}

func (u usecase) CheckOtp(data *domainEntity.CheckOtpData) error {
	defaultErr := fmt.Errorf("OTP provided is not valid. Please check the code and try again")
	phoneNumber, err := stringHelper.FormatIDNPhoneNumber(data.WhatsappNumber)
	if err != nil {
		return err
	}

	if u.repository.MaxCheckOtpAttempt(phoneNumber) {
		// return error when validation attempt has been performed 3 times
		return defaultErr
	}

	otpHash, _, err := u.repository.GetOtp(phoneNumber)
	if err != nil {
		u.manager.Log().Error(err)
		return nil
	}

	if otpHash == "" {
		return defaultErr
	}

	if !passwordHelper.CheckPassword(otpHash, u.makePassword(data.Otp, data.Token)) {
		if err := u.repository.OtpErrorCounter(phoneNumber); err != nil {
			u.manager.Log().Error(err)
			return nil
		}

		if u.repository.MaxCheckOtpAttempt(phoneNumber) {
			// unset otp on redis
			u.repository.UnsetOtp(phoneNumber)
		}

		return defaultErr
	}

	data.WhatsappNumber = phoneNumber
	data.Valid = true

	u.repository.UnsetOtp(phoneNumber)
	return nil
}

func (u usecase) makePassword(otp string, token string) string {
	return fmt.Sprintf("%s::%s", otp, token)
}

func (u usecase) otpHashing(otp string, token string) string {
	hashedPassword, _ := passwordHelper.MakePassword(u.makePassword(otp, token))
	return hashedPassword
}

func (u usecase) sendOtpToWhatsapp(phoneNumber string, otp string) {
	if u.manager.Config().Debug {
		u.printOtpToLog(phoneNumber, otp)
		return
	}

	whatsapp := whatsappHelper.New(u.manager.Config().MetaBusinessID, u.manager.Config().MetaAccessToekn)
	if err := whatsapp.SendOtp(phoneNumber, otp, strconv.Itoa(int(domainConst.OtpRequestExpiredMinute))); err != nil {
		u.manager.Log().Error(err)
		return
	}

	u.manager.Log().Infof("OTP code successfully sent to %s", phoneNumber)
}

func (u usecase) printOtpToLog(phoneNumber string, otp string) {
	u.manager.Log().WithFields(logrus.Fields{
		"PhoneNumber": phoneNumber,
		"OTP":         otp,
	}).Info("[DEBUG Mode] OTP Code")
}
