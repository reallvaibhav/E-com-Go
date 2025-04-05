package domain

type Order struct {
	ID       string         `json:"id"`
	UserID   string         `json:"user_id"`
	Products map[string]int `json:"products"` // product_id:quantity
	Status   string         `json:"status"`   // "pending|completed|cancelled"
	Total    float64        `json:"total"`
}
