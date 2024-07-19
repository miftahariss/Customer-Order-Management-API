// controller/customer_controller.go
package controller

import (
	"kstyle-hub/model"
	"kstyle-hub/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CustomerController struct {
	Service *service.CustomerService
}

func NewCustomerController(service *service.CustomerService) *CustomerController {
	return &CustomerController{Service: service}
}

// @Summary Get customers with pagination
// @Description Get customers with pagination
// @Accept  json
// @Produce  json
// @Param page query int false "Page number" default(1)
// @Param size query int false "Page size" default(10)
// @Success 200 {array} model.Customers
// @Router /customers [get]
func (c *CustomerController) GetCustomers(ctx echo.Context) error {
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
	offset, _ := strconv.Atoi(ctx.QueryParam("offset"))

	customers, err := c.Service.GetAllCustomers(limit, offset)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, customers)
}

// @Summary Get customer by ID
// @Description Get customer by ID
// @Accept  json
// @Produce  json
// @Param id path int true "Customer ID"
// @Success 200 {object} model.Customers
// @Failure 404 {object} string "Customer not found"
// @Router /customers/{id} [get]
func (c *CustomerController) GetCustomer(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	customer, err := c.Service.GetCustomerByID(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err.Error())
	}
	return ctx.JSON(http.StatusOK, customer)
}

// @Summary Create a new customer
// @Description Create a new customer
// @Accept  json
// @Produce  json
// @Param customer body model.Customers true "Customer data"
// @Success 201 {object} model.Customers
// @Failure 400 {object} string "Bad request"
// @Router /customers [post]
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

// @Summary Update an existing customer
// @Description Update an existing customer
// @Accept  json
// @Produce  json
// @Param id path int true "Customer ID"
// @Param customer body model.Customers true "Customer data"
// @Success 200 {object} model.Customers
// @Failure 400 {object} string "Bad request"
// @Failure 404 {object} string "Customer not found"
// @Router /customers/{id} [put]
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

// @Summary Delete a customer
// @Description Delete a customer
// @Accept  json
// @Produce  json
// @Param id path int true "Customer ID"
// @Success 204 {object} string "No Content"
// @Failure 404 {object} string "Customer not found"
// @Router /customers/{id} [delete]
func (c *CustomerController) DeleteCustomer(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := c.Service.DeleteCustomer(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, "Customer deleted successfully")
}

// @Summary Search customers
// @Description Search customers
// @Accept  json
// @Produce  json
// @Param query query string true "Search query"
// @Success 200 {array} model.Customers
// @Router /customers/search [get]
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
