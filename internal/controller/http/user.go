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

func (c *UserController) GetUserTransactions() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var p dto.GetUserTransactionsRequest
		if err := c.val.ValidateParams(ctx, &p); err != nil {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}

		result, err := c.userService.GetUserTransactions(ctx.UserContext(), p)
		if err != nil {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}
		if len(result) == 0 {
			return ctx.SendStatus(fiber.StatusNoContent)
		}
		return ctx.JSON(NewResponse(result, nil))
	}
}

func (c *UserController) GetUsersBalances() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		result, err := c.userService.GetUsersBalances(ctx.UserContext())
		if err != nil {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}
		if len(result) == 0 {
			return ctx.SendStatus(fiber.StatusNoContent)
		}
		return ctx.JSON(NewResponse(result, nil))
	}
}

func (c *UserController) RegisterRoutes(group fiber.Router) {
	group.Post("", c.CreateUser())
	group.Get(":user_id/transactions", c.GetUserTransactions())
	group.Get("balances", c.GetUsersBalances())
}

func NewUserController(userService domain.UserService, val govalidator.Validator) *UserController {
	return &UserController{userService, val}
}
