package entity

type OrderRepositoryInterface interface {
	CreateOrder(order *Order) error
	GetOrders() ([]*Order, error)
}
