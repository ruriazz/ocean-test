package stringHelper

import (
	"encoding/base64"
	"fmt"
	"regexp"
	"strings"
	"time"

	"math/rand"
)

func FormatIDNPhoneNumber(phone string) (string, error) {
	re := regexp.MustCompile(`\D`)
	digits := re.ReplaceAllString(phone, "")

	if strings.HasPrefix(digits, "0") || strings.HasPrefix(digits, "8") {
		digits = "62" + digits[1:]
	} else if !strings.HasPrefix(digits, "62") {
		return "", fmt.Errorf("phone number must start with '0', '62', or '8'")
	}

	formatted := "+" + digits

	if len(formatted) < 13 {
		return "", fmt.Errorf("invalid phone number")
	}

	return formatted, nil
}

func GenerateOTP() string {
	rand.Seed(time.Now().UnixNano())
	otp := rand.Intn(10000)
	return fmt.Sprintf("%04d", otp)
}

func GenerateToken(length int) string {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return ""
	}

	token := base64.URLEncoding.EncodeToString(bytes)
	return token
}

func DurationToMinuteString(dur time.Duration) string {
	minutes := int(dur.Minutes())
	seconds := int(dur.Seconds()) % 60

	if minutes == 0 {
		return fmt.Sprintf("%d seconds", seconds)
	} else {
		return fmt.Sprintf("%d minutes %d seconds", minutes, seconds)
	}
}
