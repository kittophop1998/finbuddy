package http

import (
	"github.com/gin-gonic/gin"
	"github.com/kittiphop/finbuddy/internal/delivery/http/handler"
	"github.com/kittiphop/finbuddy/internal/infrastructure/db/postgres"
	"github.com/kittiphop/finbuddy/internal/usecase"
)

func WalletRouter(r *gin.Engine) {
	walletRepo := postgres.NewWalletRepository()
	walletUsecase := usecase.NewGetWalletsUseCase(walletRepo)
	walletHandler := handler.NewWalletHandler(walletUsecase)

	// Define routes
	walletGroup := r.Group("/api/v1")
	walletGroup.GET("/wallet/wallets", walletHandler.GetWallets)
}
