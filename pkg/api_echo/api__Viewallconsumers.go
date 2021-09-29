package api_echo

import (
	model "consumerinformationmodule/pkg/model"
	arfviewallconsumers "consumerinformationmodule/pkg/rule/ARF_ViewAllConsumers"
	util "consumerinformationmodule/pkg/util"
	"context"
	"reflect"
	"strconv"
)

import (
	"github.com/labstack/echo/v4"
)

// @Summary ViewAllConsumers
// @Tags consumerinformationmodule
// @Accept json
// @Produce json
// @Success 200 {object} controller_echo.ListConsumerResponseViewAllConsumers
// @Router /consumerinformationmodule/api/consumer [get]
func ViewAllConsumers(c echo.Context) ([]*model.Consumer, error) {

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
	result, err := arfviewallconsumers.ARF_ViewAllConsumers(&cntxt, config)
	if reflect.TypeOf(result) == reflect.TypeOf([]*model.Consumer{}) || reflect.TypeOf(result) == reflect.TypeOf([]*model.Consumer{}) {
		return result.([]*model.Consumer), err
	} else {
		return nil, err
	}
}
