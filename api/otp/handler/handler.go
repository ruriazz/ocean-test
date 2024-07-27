package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	domainEntity "github.com/ruriazz/ocean-test/api/otp/domain/entity"
	domainInterface "github.com/ruriazz/ocean-test/api/otp/domain/interface"
	"github.com/ruriazz/ocean-test/api/otp/usecase"
	"github.com/ruriazz/ocean-test/pkg/manager"
	apiUtil "github.com/ruriazz/ocean-test/utils/api"
)

type handler struct {
	manager  manager.Manager
	usecase  domainInterface.OtpUsecase
	validate *validator.Validate
}

func New(manager manager.Manager) domainInterface.OtpHandler {
	return &handler{
		manager:  manager,
		usecase:  usecase.New(manager),
		validate: validator.New(),
	}
}

func (h handler) CreateOtp(ctx *fiber.Ctx) error {
	var data = new(domainEntity.CreateOtpData)
	if err := ctx.BodyParser(data); err != nil {
		return apiUtil.SendJson(ctx, apiUtil.Field{
			Status:  apiUtil.StatusUnprocessableEntity,
			Message: "Cannot parse JSON",
			Errors:  err.Error(),
		})
	}

	if err := h.validate.Struct(data); err != nil {
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, err.StructField()+" "+err.Tag())
		}
		return apiUtil.SendJson(ctx, apiUtil.Field{
			Status:  apiUtil.StatusBadRequest,
			Message: "Validation error",
			Errors:  errors,
		})
	}

	if err := h.usecase.CreateOtp(apiUtil.RequestID(ctx), data); err != nil {
		return apiUtil.SendJson(ctx, apiUtil.Field{
			Status:  apiUtil.StatusBadRequest,
			Message: "Create otp failed",
			Errors:  err.Error(),
		})
	}

	return apiUtil.SendJson(ctx, apiUtil.Field{
		Message: "OTP created",
		Data:    data,
	})
}

func (h handler) CheckOtp(ctx *fiber.Ctx) error {
	var data = new(domainEntity.CheckOtpData)
	data.WhatsappNumber = ctx.Params("phoneNumber")

	if err := ctx.BodyParser(data); err != nil {
		return apiUtil.SendJson(ctx, apiUtil.Field{
			Status:  apiUtil.StatusUnprocessableEntity,
			Message: "Cannot parse JSON",
			Errors:  err.Error(),
		})
	}

	if err := h.validate.Struct(data); err != nil {
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, err.StructField()+" "+err.Tag())
		}
		return apiUtil.SendJson(ctx, apiUtil.Field{
			Status:  apiUtil.StatusBadRequest,
			Message: "Validation error",
			Errors:  errors,
		})
	}

	if err := h.usecase.CheckOtp(data); err != nil {
		return apiUtil.SendJson(ctx, apiUtil.Field{
			Status:  apiUtil.StatusBadRequest,
			Message: "Invalid OTP",
			Errors:  err.Error(),
		})
	}

	return apiUtil.SendJson(ctx, apiUtil.Field{
		Data: data,
	})
}
