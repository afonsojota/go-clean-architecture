package graph

import (
	"go-clean-architecture/internal/usecase/order"
)

type Resolver struct {
	CreateOrderUseCase order.CreateOrderUseCase
	ListOrderUseCase   order.ListOrderUseCase
}
