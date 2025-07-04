package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kittiphop/finbuddy/internal/usecase"
)

type WalletHandler struct {
	usecase *usecase.GetWalletsUseCase
}

func NewWalletHandler(usecase *usecase.GetWalletsUseCase) *WalletHandler {
	return &WalletHandler{usecase: usecase}
}

func (wh *WalletHandler) GetWallets(ctx *gin.Context) {
	wallets, err := wh.usecase.GetWallets()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve wallets"})
		return
	}

	ctx.JSON(http.StatusOK, wallets)
}
