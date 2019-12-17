package domain

type Item struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Amount      int64  `json:"amount"`
}

type Items []Item
