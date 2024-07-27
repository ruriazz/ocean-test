package api

import (
	otp "github.com/ruriazz/ocean-test/api/otp/domain"
	"github.com/ruriazz/ocean-test/pkg/manager"
)

func Init(manager manager.Manager) {
	otp.LoadRouter(manager)
}
