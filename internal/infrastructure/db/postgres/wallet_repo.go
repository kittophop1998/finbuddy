package postgres

import "github.com/kittiphop/finbuddy/internal/domain/model"

type PostgresWalletRepository struct{}

func NewWalletRepository() *PostgresWalletRepository {
	return &PostgresWalletRepository{}
}

func (pw PostgresWalletRepository) GetWallets() (model.Wallets, error) {
	wallets := model.Wallets{
		{
			ID:           "1",
			Name:         "My Wallet",
			LoanName:     "Personal Loan",
			TotalLoan:    "10000.00",
			TotalPayment: "1000.00",
			Status:       "PAY",
			UserID:       "123",
			CreatedAt:    "2023-10-01T12:00:00Z",
		},
		{
			ID:           "2",
			Name:         "Savings Wallet",
			LoanName:     "Home Loan",
			TotalLoan:    "50000.00",
			TotalPayment: "5000.00",
			Status:       "PAID",
			UserID:       "123",
			CreatedAt:    "2023-10-02T12:00:00Z",
		},
	}

	return wallets, nil
}
