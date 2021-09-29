package controller_echo

import (
	apiecho "consumerinformationmodule/pkg/api_echo"
	model "consumerinformationmodule/pkg/model"
	"net/http"
)

import (
	"github.com/labstack/echo/v4"
)

// @title consumerinformationmodule API Documentation
// @version 1.0

func CreateConsumer(c echo.Context) error {

	result, err := apiecho.CreateConsumer(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}

func DeleteConsumers(c echo.Context) error {

	result, err := apiecho.DeleteConsumers(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}

func EditConsumer(c echo.Context) error {

	result, err := apiecho.EditConsumer(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}

type ListConsumerResponseFilterCompanyById struct {
	Consumer []*model.Consumer `json:"consumer"`
}

func FilterCompanyById(c echo.Context) error {

	result, err := apiecho.FilterCompanyById(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		value := ListConsumerResponseFilterCompanyById{}
		value.Consumer = result
		return c.JSON(http.StatusOK, value)
	}
}

type ListConsumerResponseFindConsumer struct {
	Consumer []*model.Consumer `json:"consumer"`
}

func FindConsumer(c echo.Context) error {

	result, err := apiecho.FindConsumer(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		value := ListConsumerResponseFindConsumer{}
		value.Consumer = result
		return c.JSON(http.StatusOK, value)
	}
}

type ListConsumerResponseViewAllConsumers struct {
	Consumer []*model.Consumer `json:"consumer"`
}

func ViewAllConsumers(c echo.Context) error {

	result, err := apiecho.ViewAllConsumers(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		value := ListConsumerResponseViewAllConsumers{}
		value.Consumer = result
		return c.JSON(http.StatusOK, value)
	}
}
