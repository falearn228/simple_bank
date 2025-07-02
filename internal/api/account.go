package api

import (
	"database/sql"
	"net/http"

	db "bobbabank/internal/db/sqlc"

	"github.com/gin-gonic/gin"
)

type createAccountReq struct {
	Owner string `json:"owner" binding:"required"`
	// При создании нулевого аккаунта, баланс всегда должен быть равен нулю
	// Balance  int64  `json:"balance" binding:"required"`
	Currency string `json:"currency" binding:"required,currency"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResp(err))
		return
	}

	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}

	acc, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResp(err))
		return
	}

	ctx.JSON(http.StatusOK, acc)
}

type getAccountReq struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getAccount(ctx *gin.Context) {
	var req getAccountReq
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResp(err))
		return
	}

	acc, err := server.store.GetAccount(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResp(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResp(err))
		return
	}

	ctx.JSON(http.StatusOK, acc)
}

type listAccountReq struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listAccount(ctx *gin.Context) {
	var req listAccountReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResp(err))
		return
	}

	arg := db.ListAccountsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	acc, err := server.store.ListAccounts(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResp(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResp(err))
		return
	}

	ctx.JSON(http.StatusOK, acc)
}

func (server *Server) deleteAccount(ctx *gin.Context) {
	var req getAccountReq
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResp(err))
		return
	}

	err := server.store.DeleteAccount(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResp(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResp(err))
		return
	}

	ctx.String(http.StatusOK, "Account was deleted!")
}

type updateAccountReq struct {
	ID      int64 `form:"id" binding:"required,min=1"`
	Balance int64 `form:"balance" binding:"required,min=0"`
}

func (server *Server) updateAccount(ctx *gin.Context) {
	var req updateAccountReq
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResp(err))
		return
	}

	arg := db.UpdateAccountParams{
		ID:      req.ID,
		Balance: req.Balance,
	}

	acc, err := server.store.UpdateAccount(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResp(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResp(err))
		return
	}

	ctx.JSON(http.StatusOK, acc)
}
