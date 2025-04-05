package usecase

import (
	"order-service/internal/order"
	"time"

	"github.com/google/uuid"
)

type OrderRepo interface {
	Save(order.Order) error
	FindByID(id string) (order.Order, bool)
	Update(id string, o order.Order) error
	List() []order.Order
}

type OrderUsecase struct {
	repo OrderRepo
}

func NewOrderUsecase(r OrderRepo) *OrderUsecase {
	return &OrderUsecase{repo: r}
}

func (u *OrderUsecase) Create(productID string, quantity int, price float64) (order.Order, error) {
	o := order.Order{
		ID:         uuid.New().String(),
		ProductID:  productID,
		Quantity:   quantity,
		TotalPrice: float64(quantity) * price,
		Status:     "pending",
		CreatedAt:  time.Now(),
	}
	u.repo.Save(o)
	return o, nil
}

func (u *OrderUsecase) GetByID(id string) (order.Order, bool) {
	return u.repo.FindByID(id)
}

func (u *OrderUsecase) List() []order.Order {
	return u.repo.List()
}

func (u *OrderUsecase) UpdateStatus(id, status string) (order.Order, bool) {
	o, found := u.repo.FindByID(id)
	if !found {
		return order.Order{}, false
	}
	o.Status = status
	u.repo.Update(id, o)
	return o, true
}
