package domain

import (
	"github.com/ruriazz/ocean-test/api/otp/handler"
	"github.com/ruriazz/ocean-test/pkg/manager"
)

func LoadRouter(manager manager.Manager) {
	app := manager.Server().App()
	handler := handler.New(manager)

	app.Post("/otp", handler.CreateOtp)
	app.Post("/otp/:phoneNumber", handler.CheckOtp)
}
