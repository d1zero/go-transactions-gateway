package http

import (
	"github.com/gofiber/fiber/v2"
	"go-transactions-gateway/internal/domain"
	"go-transactions-gateway/internal/domain/dto"
	"go-transactions-gateway/pkg/govalidator"
)

type UserController struct {
	userService domain.UserService
	val         govalidator.Validator
}

func (c *UserController) CreateUser() fiber.Handler {
	return func(ctx *fiber.Ctx) (err error) {
		var p dto.CreateUser
		if err = c.val.ValidateRequestBody(ctx, &p); err != nil {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}
		result, err := c.userService.CreateUser(ctx.UserContext(), p)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(NewErrResponse(err))
		}
		return ctx.Status(fiber.StatusCreated).JSON(NewResponse(result, nil))
	}
}

func (c *UserController) RegisterRoutes(group fiber.Router) {
	group.Post("", c.CreateUser())
}

func NewUserController(userService domain.UserService, val govalidator.Validator) *UserController {
	return &UserController{userService, val}
}
