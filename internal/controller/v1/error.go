package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type response struct {
	Error string `json:"error" example:"message"`
}

var (
	ErrIllegalCurrencyCode    = errors.New("ErrIllegalCurrencyCode")
	ErrNegativeDepositCode    = errors.New("ErrNegativeDepositCode")
	ErrNegativeWithdrawalCode = errors.New("ErrNegativeWithdrawalCode")
	//ErrSpendingBudgetExceeded = errors.New("ErrSpendingBudgetExceeded")
)

func errorResponse(c *gin.Context, err error) {

	c.AbortWithStatusJSON(http.StatusBadRequest, response{err.Error()})

}
