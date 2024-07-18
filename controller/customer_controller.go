// controller/customer_controller.go
package controller

import (
    "net/http"
    "kstyle-hub/model"
    "kstyle-hub/service"
    "strconv"

    "github.com/labstack/echo/v4"
)

type CustomerController struct {
    Service *service.CustomerService
}

func NewCustomerController(service *service.CustomerService) *CustomerController {
    return &CustomerController{Service: service}
}

func (c *CustomerController) GetCustomers(ctx echo.Context) error {
    limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
    offset, _ := strconv.Atoi(ctx.QueryParam("offset"))

    customers, err := c.Service.GetAllCustomers(limit, offset)
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, err.Error())
    }
    return ctx.JSON(http.StatusOK, customers)
}

func (c *CustomerController) GetCustomer(ctx echo.Context) error {
    id, _ := strconv.Atoi(ctx.Param("id"))

    customer, err := c.Service.GetCustomerByID(uint(id))
    if err != nil {
        return ctx.JSON(http.StatusNotFound, err.Error())
    }
    return ctx.JSON(http.StatusOK, customer)
}

func (c *CustomerController) CreateCustomer(ctx echo.Context) error {
    customer := new(model.Customer)
    if err := ctx.Bind(customer); err != nil {
        return ctx.JSON(http.StatusBadRequest, err.Error())
    }

    err := c.Service.CreateCustomer(customer)
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, err.Error())
    }
    return ctx.JSON(http.StatusCreated, customer)
}

func (c *CustomerController) UpdateCustomer(ctx echo.Context) error {
    id, _ := strconv.Atoi(ctx.Param("id"))
    customer := new(model.Customer)
    if err := ctx.Bind(customer); err != nil {
        return ctx.JSON(http.StatusBadRequest, err.Error())
    }
    customer.ID = uint(id)

    err := c.Service.UpdateCustomer(customer)
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, err.Error())
    }
    return ctx.JSON(http.StatusOK, customer)
}

func (c *CustomerController) DeleteCustomer(ctx echo.Context) error {
    id, _ := strconv.Atoi(ctx.Param("id"))

    err := c.Service.DeleteCustomer(uint(id))
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, err.Error())
    }
    return ctx.JSON(http.StatusOK, "Customer deleted successfully")
}

func (c *CustomerController) SearchCustomers(ctx echo.Context) error {
    query := ctx.QueryParam("query")
    limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
    offset, _ := strconv.Atoi(ctx.QueryParam("offset"))

    customers, err := c.Service.SearchCustomers(query, limit, offset)
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, err.Error())
    }
    return ctx.JSON(http.StatusOK, customers)
}
