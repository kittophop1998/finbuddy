package usecase

import (
	"github.com/kittiphop/finbuddy/internal/domain/model"
	"github.com/kittiphop/finbuddy/internal/domain/port"
)

type GetWalletsUseCase struct {
	repo port.WalletRepository
}

func NewGetWalletsUseCase(repo port.WalletRepository) *GetWalletsUseCase {
	return &GetWalletsUseCase{repo: repo}
}

func (gw GetWalletsUseCase) GetWallets() (model.Wallets, error) {

	return gw.repo.GetWallets()
}
