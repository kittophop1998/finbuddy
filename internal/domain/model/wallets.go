package model

type Wallet struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	LoanName     string `json:"loanName"`
	TotalLoan    string `json:"totalLoan"`
	TotalPayment string `json:"totalPayment"`
	Status       string `json:"status"` // e.g., "PAY", "PAID"
	UserID       string `json:"user_id"`
	CreatedAt    string `json:"created_at"`
}

type Wallets []Wallet
