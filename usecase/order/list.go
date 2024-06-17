package order

import (
	"go-clean-architecture/internal/entity"
)

type OrdersOutputDTO struct {
	Orders []OrderDTO `json:"orders"`
}

type OrderDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *ListOrderUseCase) GetOrders() (OrdersOutputDTO, error) {
	orders, err := c.OrderRepository.GetOrders()
	if err != nil {
		return OrdersOutputDTO{}, err
	}

	ordersDTO := make([]OrderDTO, 0)
	for _, order := range orders {
		ordersDTO = append(ordersDTO, OrderDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}

	return OrdersOutputDTO{
		Orders: ordersDTO,
	}, nil
}
