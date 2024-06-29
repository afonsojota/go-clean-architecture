// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"database/sql"
	"github.com/google/wire"
	"github.com/afonsojota/go-clean-architecture/internal/entity"
	"github.com/afonsojota/go-clean-architecture/internal/event"
	"github.com/afonsojota/go-clean-architecture/internal/infra/database"
	"github.com/afonsojota/go-clean-architecture/internal/infra/web"
	"github.com/afonsojota/go-clean-architecture/internal/usecase"
	"github.com/afonsojota/go-clean-architecture/pkg/events"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from wire.go:

func NewCreateOrderUseCase(db *sql.DB, eventDispatcherInterface events.EventDispatcherInterface) *usecase.CreateOrderUseCase {
	orderRepository := database.NewOrderRepository(db)
	orderCreated := event.NewOrderCreated()
	createOrderUseCase := usecase.NewCreateOrderUseCase(orderRepository, orderCreated, eventDispatcherInterface)
	return createOrderUseCase
}

func NewListOrdersUseCase(db *sql.DB) *usecase.ListOrdersUseCase {
	orderRepository := database.NewOrderRepository(db)
	listOrdersUseCase := usecase.NewListOrdersUseCase(orderRepository)
	return listOrdersUseCase
}

func NewWebOrderHandler(db *sql.DB, eventDispatcherInterface events.EventDispatcherInterface) *web.WebOrderHandler {
	orderRepository := database.NewOrderRepository(db)
	orderCreated := event.NewOrderCreated()
	webOrderHandler := web.NewWebOrderHandler(eventDispatcherInterface, orderRepository, orderCreated)
	return webOrderHandler
}

// wire.go:

var setOrderRepositoryDependency = wire.NewSet(database.NewOrderRepository, wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)))

var setEventDispatcherDependency = wire.NewSet(events.NewEventDispatcher, event.NewOrderCreated, wire.Bind(new(events.EventInterface), new(*event.OrderCreated)), wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)))

var setOrderCreatedEvent = wire.NewSet(event.NewOrderCreated, wire.Bind(new(events.EventInterface), new(*event.OrderCreated)))
