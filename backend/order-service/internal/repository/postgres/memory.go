package repository

import (
	"order-service/internal/order"
	"sync"
)

type InMemoryOrderRepo struct {
	orders map[string]order.Order
	mu     sync.RWMutex
}

func NewInMemoryOrderRepo() *InMemoryOrderRepo {
	return &InMemoryOrderRepo{
		orders: make(map[string]order.Order),
	}
}

func (r *InMemoryOrderRepo) Save(o order.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.orders[o.ID] = o
	return nil
}

func (r *InMemoryOrderRepo) FindByID(id string) (order.Order, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	o, exists := r.orders[id]
	return o, exists
}

func (r *InMemoryOrderRepo) Update(id string, updated order.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.orders[id] = updated
	return nil
}

func (r *InMemoryOrderRepo) List() []order.Order {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var list []order.Order
	for _, o := range r.orders {
		list = append(list, o)
	}
	return list
}
