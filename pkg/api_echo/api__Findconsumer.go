package api_echo

import (
	model "consumerinformationmodule/pkg/model"
	arffindconsumer "consumerinformationmodule/pkg/rule/ARF_FindConsumer"
	util "consumerinformationmodule/pkg/util"
	"context"
	"reflect"
	"strconv"
)

import (
	"github.com/labstack/echo/v4"
)

// @Summary FindConsumer
// @Tags consumerinformationmodule
// @Accept json
// @Produce json
// @Param consumerid query string true  " string consumerid query "
// @Success 200 {object} controller_echo.ListConsumerResponseFindConsumer
// @Router /consumerinformationmodule/api/consumerbyid [get]
func FindConsumer(c echo.Context) ([]*model.Consumer, error) {

	consumerid := c.QueryParam("consumerid")
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
	result, err := arffindconsumer.ARF_FindConsumer(consumerid, &cntxt, config)
	if reflect.TypeOf(result) == reflect.TypeOf([]*model.Consumer{}) || reflect.TypeOf(result) == reflect.TypeOf([]*model.Consumer{}) {
		return result.([]*model.Consumer), err
	} else {
		return nil, err
	}
}
