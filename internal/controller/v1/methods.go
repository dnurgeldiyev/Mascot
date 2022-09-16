package v1

import (
	"dovran/mascot/internal/entity"
	"dovran/mascot/internal/usecase"
	"dovran/mascot/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	handler.POST("/add-player-balance", r.addPlayerWithBalance)
	handler.POST("/get-balance", r.getBalance)
	handler.POST("/with-draw-deposit", r.withDrawAndDeposit)
	handler.POST("/rollback-transaction", r.rollBackTransaction)

}

func (r *routes) addPlayerWithBalance(c *gin.Context) {

	var playerWithBalance PlayerWithBalance
	var err error

	if err = c.ShouldBindJSON(&playerWithBalance); err != nil {
		r.l.Error(err, "http-v1-addPlayerWithBalance")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	if err = playerWithBalance.isValid(); err != nil {
		r.l.Error(err, "http-v1-playerWithBalance.isValid")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	err = r.u.AddPlayerWithBalance(playerWithBalance.PlayerName, playerWithBalance.Balance)
	if err != nil {
		r.l.Error(err, "http-v1-getBalance r.u.GetBalance")
		errorResponse(c, http.StatusBadRequest, "player not found")
		return
	}

	c.JSON(http.StatusOK, `{}`)

}
func (r *routes) getBalance(c *gin.Context) {

	var balance Balance
	var err error

	if err = c.ShouldBindJSON(&balance); err != nil {
		r.l.Error(err, "http-v1-getBalance")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	if err = balance.isValid(); err != nil {
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

	c.JSON(http.StatusOK, BalanceResponse{
		Balance:        getBalance.Balance,
		FreeRoundsLeft: getBalance.FreeRoundsLeft,
	})

}
func (r *routes) withDrawAndDeposit(c *gin.Context) {

	var drawAndDeposit DrawAndDeposit
	var err error

	if err = c.ShouldBindJSON(&drawAndDeposit); err != nil {
		r.l.Error(err, "http-v1-withDrawAndDeposit")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	if err = drawAndDeposit.isValid(); err != nil {
		r.l.Error(err, "http-v1-drawAndDeposit.isValid")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	dtoDrawAndDeposit := entity.DrawAndDepositDTO{
		CallerId:             drawAndDeposit.CallerId,
		PlayerName:           drawAndDeposit.PlayerName,
		WithDraw:             drawAndDeposit.WithDraw,
		Deposit:              drawAndDeposit.Deposit,
		Currency:             string(drawAndDeposit.Currency),
		TransactionRef:       drawAndDeposit.TransactionRef,
		GameRoundRef:         drawAndDeposit.GameRoundRef,
		GameId:               drawAndDeposit.GameId,
		Source:               drawAndDeposit.Source,
		Reason:               string(drawAndDeposit.Reason),
		SessionId:            drawAndDeposit.SessionId,
		SessionAlternativeId: drawAndDeposit.SessionAlternativeId,
		BonusId:              drawAndDeposit.BonusId,
		ChargeFreeRounds:     drawAndDeposit.ChargeFreeRounds,
	}

	resDrawAndDeposit, err := r.u.WithDrawAndDeposit(dtoDrawAndDeposit)
	if err != nil {
		r.l.Error(err, "http-v1-getBalance r.u.WithDrawAndDeposit")
		errorResponse(c, http.StatusBadRequest, "invalid request")
		return
	}

	c.JSON(http.StatusOK, DrawAndDepositResponse{
		NewBalance:     resDrawAndDeposit.Balance,
		TransactionId:  uuid.NewString(),
		FreeRoundsLeft: resDrawAndDeposit.FreeRoundsLeft,
	})

}
func (r *routes) rollBackTransaction(c *gin.Context) {

	var rollBackTransaction RollBackTransaction
	var err error

	if err = c.ShouldBindJSON(&rollBackTransaction); err != nil {
		r.l.Error(err, "http-v1-rollBackTransaction")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	if err = rollBackTransaction.isValid(); err != nil {
		r.l.Error(err, "http-v1-rollBackTransaction.isValid")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}
	dtoRollBack := entity.RollBackTransactionDTO{
		CallerId:             rollBackTransaction.CallerId,
		PlayerName:           rollBackTransaction.PlayerName,
		TransactionRef:       rollBackTransaction.TransactionRef,
		GameId:               rollBackTransaction.GameId,
		SessionId:            rollBackTransaction.SessionId,
		SessionAlternativeId: rollBackTransaction.SessionAlternativeId,
		RoundId:              rollBackTransaction.RoundId,
	}
	err = r.u.RollBackTransaction(dtoRollBack)
	if err != nil {
		r.l.Error(err, "http-v1-getBalance r.u.RollBackTransaction")
		errorResponse(c, http.StatusBadRequest, "invalid request")
		return
	}

	c.JSON(http.StatusOK, `{}`)
}
