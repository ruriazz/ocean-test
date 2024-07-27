package repository

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	domainConst "github.com/ruriazz/ocean-test/api/otp/domain/const"
	domainInterface "github.com/ruriazz/ocean-test/api/otp/domain/interface"
	"github.com/ruriazz/ocean-test/pkg/manager"
)

type repository struct {
	manager manager.Manager
}

func New(manager manager.Manager) domainInterface.OtpRepository {
	return repository{
		manager: manager,
	}
}

func (r repository) NextCreateOtpTime(phoneNumber string) (time.Duration, error) {
	keys, err := r.manager.Redis().Keys(r.makeOtpKey(phoneNumber, ""))
	if err != nil {
		return 0, err
	}

	if len(keys) < 3 {
		return 0, nil
	}

	ttl, err := r.manager.Redis().TTL(keys[len(keys)-1])
	if err != nil {
		return 0, nil
	}

	return ttl, nil
}

func (r repository) SaveOtp(requestid string, phoneNumber string, value string) error {
	r.UnsetOtp(phoneNumber)

	value = fmt.Sprintf("%s||%d", value, time.Now().Unix())

	return r.manager.Redis().SetString(r.makeOtpKey(phoneNumber, requestid), value, domainConst.OtpRequestExpiredMinute*time.Minute)
}

func (r repository) UnsetOtp(phoneNumber string) {
	keys, _ := r.manager.Redis().Keys(r.makeOtpKey(phoneNumber, ""))
	for _, key := range keys {
		r.manager.Redis().SetString(key, "", domainConst.OtpRequestExpiredMinute*time.Minute)
	}
}

func (r repository) OtpErrorCounter(phoneNumber string) error {
	var count int = 0
	var key string = r.makeOtpFailKey(phoneNumber)

	if val, _ := r.manager.Redis().GetString(key); val != "" {
		count, _ = strconv.Atoi(val)
	}

	count += 1
	return r.manager.Redis().SetString(key, strconv.Itoa(count), domainConst.OtpRequestExpiredMinute*time.Minute)
}

func (r repository) MaxCheckOtpAttempt(phoneNumber string) bool {
	key := r.makeOtpFailKey(phoneNumber)
	if val, _ := r.manager.Redis().GetString(key); val != "" {
		count, _ := strconv.Atoi(val)
		if count >= 3 {
			return true
		}
	}

	return false
}

func (r repository) GetOtp(phoneNumber string) (string, time.Time, error) {
	var value string
	var creationTime time.Time

	keys, err := r.manager.Redis().Keys(r.makeOtpKey(phoneNumber, ""))
	if err != nil {
		return "", time.Time{}, err
	}

	for _, key := range keys {
		val, err := r.manager.Redis().GetString(key)
		if err != nil {
			return "", time.Time{}, err
		}

		if val != "" {
			parts := strings.Split(val, "||")
			if len(parts) == 2 {
				creationTimeUnix, err := strconv.ParseInt(parts[1], 10, 64)
				if err != nil {
					return "", time.Time{}, err
				}

				value = parts[0]
				creationTime = time.Unix(creationTimeUnix, 0)
			}
			break
		}
	}

	return value, creationTime, nil
}

func (r repository) makeOtpKey(phoneNumber string, requestid string) string {
	return fmt.Sprintf("OTP%s::%s", phoneNumber, requestid)
}

func (r repository) makeOtpFailKey(phoneNumber string) string {
	return fmt.Sprintf("OTP-Fail_%s", phoneNumber)
}
