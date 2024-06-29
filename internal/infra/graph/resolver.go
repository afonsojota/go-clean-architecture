package graph

import "github.com/afonsojota/go-clean-architecture/internal/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	usecase.CreateOrderUseCase
	usecase.ListOrdersUseCase
}
