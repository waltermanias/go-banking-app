package dto

type AccountRequest struct {
	CustomerId  string  `json:"customerId"`
	OpeningDate string  `json:"openingDate"`
	AccountType string  `json:"accountType"`
	Amount      float64 `json:"amount"`
	Status      string  `json:"status"`
}

type AccountResponse struct {
	AccountId   string  `json:"id"`
	CustomerId  string  `json:"customerId"`
	OpeningDate string  `json:"openingDate"`
	AccountType string  `json:"accountType"`
	Amount      float64 `json:"amount"`
	Status      string  `json:"status"`
}

func (a AccountRequest) StatusToString() string {
	if a.Status == "active" {
		return "1"
	} else if a.Status == "inactive" {
		return "0"
	}
	return "0"
}
