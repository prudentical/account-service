package api

import (
	"account-service/internal/model"
	"account-service/internal/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AccountHandler interface {
	GetById(c echo.Context) error
	GetAll(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
	HandleRoutes(e *echo.Echo)
}

type AccountHandlerImpl struct {
	service service.AccountService
}

func NewAccountHandler(service service.AccountService) AccountHandler {
	return AccountHandlerImpl{service}
}

func (h AccountHandlerImpl) HandleRoutes(e *echo.Echo) {
	e.GET("/users/:user_id/accounts", h.GetAll)
	e.POST("/users/:user_id/accounts", h.Create)
	e.GET("/users/:user_id/accounts/:id", h.GetById)
	e.PUT("/users/:user_id/accounts/:id", h.Update)
	e.DELETE("/users/:user_id/accounts/:id", h.Delete)
}

func (h AccountHandlerImpl) GetById(c echo.Context) error {
	userIdStr := c.Param("user_id")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		return InvalidIDError{TypeName: "User", Id: userIdStr}
	}

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return InvalidIDError{Type: model.Account{}, Id: idStr}
	}

	account, err := h.service.GetById(userId, id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, account)
}

func (h AccountHandlerImpl) GetAll(c echo.Context) error {
	userIdStr := c.Param("user_id")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		return InvalidIDError{TypeName: "User", Id: userIdStr}
	}

	pageStr := c.QueryParam("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	sizeStr := c.QueryParam("size")
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		size = 20
	}

	accounts, err := h.service.GetAll(userId, page, size)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, accounts)
}

func (h AccountHandlerImpl) Create(c echo.Context) error {
	userIdStr := c.Param("user_id")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		return InvalidIDError{TypeName: "User", Id: userIdStr}
	}

	var account model.Account
	if err := c.Bind(&account); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(account); err != nil {
		return err
	}

	createdAccount, err := h.service.Create(userId, account)
	if err != nil {
		return err
	}

	location := fmt.Sprintf("/users/%s/accounts/%d", userIdStr, createdAccount.ID)
	c.Response().Header().Set("location", location)

	return c.NoContent(http.StatusAccepted)
}

func (h AccountHandlerImpl) Update(c echo.Context) error {
	userIdStr := c.Param("user_id")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		return InvalidIDError{TypeName: "User", Id: userIdStr}
	}

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return InvalidIDError{Type: model.Account{}, Id: idStr}
	}

	var account model.Account
	if err := c.Bind(&account); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(account); err != nil {
		return err
	}

	account, err = h.service.Update(userId, id, account)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusAccepted, account)
}

func (h AccountHandlerImpl) Delete(c echo.Context) error {
	userIdStr := c.Param("user_id")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		return InvalidIDError{TypeName: "User", Id: userIdStr}
	}

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return InvalidIDError{Type: model.Account{}, Id: idStr}
	}

	err = h.service.Delete(userId, id)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}
