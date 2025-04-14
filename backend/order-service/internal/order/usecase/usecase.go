

package usecase

type Order struct {
	ID      string `json:"id"`
	UserID  string `json:"user_id"`
	Product string `json:"product"`
	Amount  int    `json:"amount"`
}

type OrderRepo interface {
	FetchAll() []Order
	FetchByID(id string) (Order, bool)
	Save(order Order)
	Remove(id string) bool
}
