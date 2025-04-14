package usecase

type Order struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	Product  string `json:"product"`
	Quantity int    `json:"quantity"`
}

type OrderUsecase interface {
	GetAll() []Order
	GetByID(id string) (Order, bool)
	Create(order Order)
	Delete(id string) bool
}

type orderUsecase struct {
	repo OrderRepo
}

func NewOrderUsecase(r OrderRepo) OrderUsecase {
	return &orderUsecase{repo: r}
}

func (uc *orderUsecase) GetAll() []Order {
	return uc.repo.FetchAll()
}

func (uc *orderUsecase) GetByID(id string) (Order, bool) {
	return uc.repo.FetchByID(id)
}

func (uc *orderUsecase) Create(order Order) {
	uc.repo.Save(order)
}

func (uc *orderUsecase) Delete(id string) bool {
	return uc.repo.Remove(id)
}

type OrderRepo interface {
	FetchAll() []Order
	FetchByID(id string) (Order, bool)
	Save(order Order)
	Remove(id string) bool
}
