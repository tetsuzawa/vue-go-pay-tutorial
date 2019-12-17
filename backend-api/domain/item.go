package domain

type Item struct {
	ID          int64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Amount      int64  `json:"amount"`
}

type Items []Item
