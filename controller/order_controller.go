// controller/order_controller.go
package controller

import (
    "net/http"
    "kstyle-hub/model"
    "kstyle-hub/service"
    "strconv"

    "github.com/labstack/echo/v4"
)

type OrderController struct {
    Service *service.OrderService
}

func NewOrderController(service *service.OrderService) *OrderController {
    return &OrderController{Service: service}
}

func (o *OrderController) GetOrders(ctx echo.Context) error {
    limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
    offset, _ := strconv.Atoi(ctx.QueryParam("offset"))

    orders, err := o.Service.GetAllOrders(limit, offset)
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, err.Error())
    }
    return ctx.JSON(http.StatusOK, orders)
}

func (o *OrderController) GetOrder(ctx echo.Context) error {
    id, _ := strconv.Atoi(ctx.Param("id"))

    order, err := o.Service.GetOrderByID(uint(id))
    if err != nil {
        return ctx.JSON(http.StatusNotFound, err.Error())
    }
    return ctx.JSON(http.StatusOK, order)
}

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

func (o *OrderController) DeleteOrder(ctx echo.Context) error {
    id, _ := strconv.Atoi(ctx.Param("id"))

    err := o.Service.DeleteOrder(uint(id))
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, err.Error())
    }
    return ctx.JSON(http.StatusOK, "Order deleted successfully")
}

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
