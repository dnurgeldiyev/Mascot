package v1

import (
	"dovran/mascot/internal/usecase"
	"dovran/mascot/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type routes struct {
	u usecase.UseCase
	l logger.Interface
}

func methods(handler *gin.RouterGroup, u usecase.UseCase, l logger.Interface) {
	r := &routes{
		u: u,
		l: l,
	}

	handler.POST("/get-balance", r.getBalance)
	handler.POST("/with-draw-deposit", r.withDrawAndDeposit)
	handler.POST("/rollback-transaction", r.rollBackTransaction)

}

func (r *routes) getBalance(c *gin.Context) {

	var balance Balance
	if err := c.ShouldBindJSON(&balance); err != nil {
		r.l.Error(err, "http-v1-getBalance")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := balance.isValid(); err != nil {
		r.l.Error(err, "http-v1-balance.isValid")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	getBalance, err := r.u.GetBalance(balance.PlayerName)
	if err != nil {
		r.l.Error(err, "http-v1-getBalance r.u.GetBalance")
		errorResponse(c, http.StatusBadRequest, "player not found")
		return
	}

	c.JSON(http.StatusOK, getBalance)

}
func (r *routes) withDrawAndDeposit(c *gin.Context) {
	var drawAndDeposit DrawAndDeposit

	if err := c.ShouldBindJSON(&drawAndDeposit); err != nil {
		r.l.Error(err, "http-v1-withDrawAndDeposit")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := drawAndDeposit.isValid(); err != nil {
		r.l.Error(err, "http-v1-drawAndDeposit.isValid")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	resDrawAndDeposit, err := r.u.WithDrawAndDeposit(drawAndDeposit)
	if err != nil {
		r.l.Error(err, "http-v1-getBalance r.u.WithDrawAndDeposit")
		errorResponse(c, http.StatusBadRequest, "invalid request")
		return
	}

	c.JSON(http.StatusOK, resDrawAndDeposit)

}
func (r *routes) rollBackTransaction(c *gin.Context) {

	var rollBackTransaction RollBackTransaction

	if err := c.ShouldBindJSON(&rollBackTransaction); err != nil {
		r.l.Error(err, "http-v1-rollBackTransaction")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := rollBackTransaction.isValid(); err != nil {
		r.l.Error(err, "http-v1-rollBackTransaction.isValid")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	err := r.u.RollBackTransaction(rollBackTransaction)
	if err != nil {
		r.l.Error(err, "http-v1-getBalance r.u.RollBackTransaction")
		errorResponse(c, http.StatusBadRequest, "invalid request")
		return
	}

	c.JSON(http.StatusOK, `{}`)
}
