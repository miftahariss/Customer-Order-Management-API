// controller/order_controller.go
package controller

import (
	"kstyle-hub/model"
	"kstyle-hub/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type OrderController struct {
	Service *service.OrderService
}

func NewOrderController(service *service.OrderService) *OrderController {
	return &OrderController{Service: service}
}

// @Summary Get orders with pagination
// @Description Get orders with pagination
// @Accept  json
// @Produce  json
// @Param page query int false "Page number" default(1)
// @Param size query int false "Page size" default(10)
// @Success 200 {array} model.Orders
// @Router /orders [get]
func (o *OrderController) GetOrders(ctx echo.Context) error {
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
	offset, _ := strconv.Atoi(ctx.QueryParam("offset"))

	orders, err := o.Service.GetAllOrders(limit, offset)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, orders)
}

// @Summary Get order by ID
// @Description Get order by ID
// @Accept  json
// @Produce  json
// @Param id path int true "Order ID"
// @Success 200 {object} model.Orders
// @Failure 404 {object} string "Order not found"
// @Router /orders/{id} [get]
func (o *OrderController) GetOrder(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	order, err := o.Service.GetOrderByID(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err.Error())
	}
	return ctx.JSON(http.StatusOK, order)
}

// @Summary Create a new order
// @Description Create a new order
// @Accept  json
// @Produce  json
// @Param order body model.Orders true "Order data"
// @Success 201 {object} model.Orders
// @Failure 400 {object} string "Bad request"
// @Router /orders [post]
func (o *OrderController) CreateOrder(ctx echo.Context) error {
	order := new(model.Order)
	if err := ctx.Bind(order); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	err := o.Service.CreateOrder(order)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusCreated, order)
}

// @Summary Update an existing order
// @Description Update an existing order
// @Accept  json
// @Produce  json
// @Param id path int true "Order ID"
// @Param order body model.Orders true "Order data"
// @Success 200 {object} model.Orders
// @Failure 400 {object} string "Bad request"
// @Failure 404 {object} string "Order not found"
// @Router /orders/{id} [put]
func (o *OrderController) UpdateOrder(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	order := new(model.Order)
	if err := ctx.Bind(order); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	order.ID = uint(id)

	err := o.Service.UpdateOrder(order)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, order)
}

// @Summary Delete an order
// @Description Delete an order
// @Accept  json
// @Produce  json
// @Param id path int true "Order ID"
// @Success 204 {object} string "No Content"
// @Failure 404 {object} string "Order not found"
// @Router /orders/{id} [delete]
func (o *OrderController) DeleteOrder(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := o.Service.DeleteOrder(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, "Order deleted successfully")
}

// @Summary Search orders
// @Description Search orders
// @Accept  json
// @Produce  json
// @Param query query string true "Search query"
// @Success 200 {array} model.Orders
// @Router /orders/search [get]
func (o *OrderController) SearchOrders(ctx echo.Context) error {
	query := ctx.QueryParam("query")
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
	offset, _ := strconv.Atoi(ctx.QueryParam("offset"))

	orders, err := o.Service.SearchOrders(query, limit, offset)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, orders)
}
