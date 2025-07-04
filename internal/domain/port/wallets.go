package port

import "github.com/kittiphop/finbuddy/internal/domain/model"

type WalletRepository interface {
	GetWallets() (model.Wallets, error)
}
