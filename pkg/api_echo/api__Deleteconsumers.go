package api_echo

import (
	arfdeleteconsumers "consumerinformationmodule/pkg/rule/ARF_DeleteConsumers"
	util "consumerinformationmodule/pkg/util"
	"context"
	"reflect"
	"strconv"
)

import (
	"github.com/labstack/echo/v4"
)

// @Summary DeleteConsumers
// @Tags consumerinformationmodule
// @Accept json
// @Produce json
// @Success 200
// @Router /consumerinformationmodule/api/consumer [delete]
func DeleteConsumers(c echo.Context) (string, error) {

	cntxt := util.CustomContext{}
	cntxt.Echo = c
	cntxt.AppGoContext = context.Background()
	config := make(map[string]interface{})
	_index, err := strconv.Atoi(c.QueryParam("_index"))
	if err != nil {
		_index = -1
	}
	config["_index"] = _index
	_size, err := strconv.Atoi(c.QueryParam("_size"))
	if err != nil {
		_size = -1
	}
	config["_size"] = _size
	result, err := arfdeleteconsumers.ARF_DeleteConsumers(&cntxt, config)
	if reflect.TypeOf(result).String() == "string" {
		return result.(string), err
	} else {
		return "", err
	}
}
