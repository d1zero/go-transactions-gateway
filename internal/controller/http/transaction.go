package http

import (
	"github.com/gofiber/fiber/v2"
	"go-transactions-gateway/internal/domain"
	"go-transactions-gateway/internal/domain/dto"
	"go-transactions-gateway/pkg/govalidator"
)

type TransactionController struct {
	val                govalidator.Validator
	transactionService domain.TransactionService
}

// GetTransactions godoc
// @Summary      GetTransactions
// @Description  Get transactions list
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Param 		 userID body dto.GetTransactionsRequest true "User ID"
// @Success      200  {object}  nil
// @Failure      400  {object}  nil
// @Router       /transaction [post]
func (c *TransactionController) GetTransactions() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var p dto.PaginationRequest
		if err := c.val.ValidateQueryParams(ctx, &p); err != nil {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}

		result, err := c.transactionService.GetTransactions(ctx.UserContext(), p)
		if err != nil {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}
		if len(result.Data) == 0 {
			return ctx.SendStatus(fiber.StatusNoContent)
		}
		return ctx.JSON(NewResponse(result.Data, result.PaginationResponse))
	}
}

func (c *TransactionController) RegisterRoutes(group fiber.Router) {
	group.Get("", c.GetTransactions())
}

func NewTransactionService(
	transactionService domain.TransactionService,
	val govalidator.Validator,
) *TransactionController {
	return &TransactionController{
		val,
		transactionService,
	}
}
