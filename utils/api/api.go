package apiUtil

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

var validate = validator.New()

func SendJson(c *fiber.Ctx, field Field) error {
	if field.Status == 0 {
		field.Status = StatusOK
	}

	return c.Status(field.Status).JSON(JsonResponse{
		Message: field.Message,
		Data:    field.Data,
		Errors:  field.Errors,
	})
}

func RequestID(c *fiber.Ctx) string {
	return c.Locals(requestid.ConfigDefault.ContextKey).(string)
}

func ParseAndValidateBody(c *fiber.Ctx, req interface{}) error {

	if err := c.BodyParser(req); err != nil {
		return SendJson(c, Field{
			Status:  StatusUnprocessableEntity,
			Message: "Cannot parse JSON",
			Errors:  err.Error(),
		})
	}

	if err := validate.Struct(req); err != nil {
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, err.StructNamespace()+" "+err.Tag())
		}
		return SendJson(c, Field{
			Status:  StatusBadRequest,
			Message: "Validation error",
			Errors:  errors,
		})
	}

	return nil
}
