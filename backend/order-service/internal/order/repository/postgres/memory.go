package repository

import (
	"sync"

	"order-service/internal/order/usecase"
)

type inMemoryRepo struct {
	data map[string]usecase.Order
	mu   sync.RWMutex
}

func NewInMemoryOrderRepo() usecase.OrderRepo {
	return &inMemoryRepo{
		data: make(map[string]usecase.Order),
	}
}

func (r *inMemoryRepo) FetchAll() []usecase.Order {
	r.mu.RLock()
	defer r.mu.RUnlock()

	orders := make([]usecase.Order, 0, len(r.data))
	for _, order := range r.data {
		orders = append(orders, order)
	}
	return orders
}

func (r *inMemoryRepo) FetchByID(id string) (usecase.Order, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	order, exists := r.data[id]
	return order, exists
}

func (r *inMemoryRepo) Save(order usecase.Order) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.data[order.ID] = order
}

func (r *inMemoryRepo) Remove(id string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.data[id]; !exists {
		return false
	}
	delete(r.data, id)
	return true
}
