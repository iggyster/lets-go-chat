package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iggyster/lets-go-chat/internal/user"
)

type RegisterData struct {
	Username string `json:"userName"`
	Password string `json:"password"`
}

type RegisterResource struct {
	Id       string `json:"id"`
	UserName string `json:"userName"`
}

func Register(ctx *fiber.Ctx) error {
	data := RegisterData{}
	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	if errors := validate(&data); errors.Count() > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}

	usr := user.New(data.Username, data.Password)
	user.Repository.Save(usr)

	return ctx.Status(fiber.StatusOK).JSON(RegisterResource{usr.Id, usr.Username})
}

func validate(data *RegisterData) Errors {
	errors := Errors{}
	if user.Repository.IsExists(data.Username) {
		errors.AddError("userName", "User already exists", "Change the username")
	}
	if data.Username == "" {
		errors.AddError("userName", "User is empty", "Change the username")
	}
	if len(data.Username) < 3 {
		errors.AddError(
			"userName",
			"User is too short",
			"Username must be greater or equal to 3 characters",
		)
	}
	if data.Password == "" {
		errors.AddError("password", "Password is empty", "Change the username")
	}
	if len(data.Password) <= 8 || len(data.Password) > 32 {
		errors.AddError(
			"password",
			"Password is invalid",
			"Password must be more then 8 and less then 32 characters long",
		)
	}
	return errors
}
