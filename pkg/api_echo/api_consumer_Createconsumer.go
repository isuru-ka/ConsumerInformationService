package api_echo

import (
	model "consumerinformationmodule/pkg/model"
	arfcreateconsumer "consumerinformationmodule/pkg/rule/ARF_CreateConsumer"
	util "consumerinformationmodule/pkg/util"
	"context"
	"reflect"
	"strconv"
)

import (
	"github.com/labstack/echo/v4"
)

// @Summary CreateConsumer
// @Tags consumerinformationmodule
// @Accept json
// @Produce json
// @Param body-param body model.Consumer true  "model.Consumer body data"
// @Success 200 {object} model.Consumer
// @Router /consumerinformationmodule/api/consumer [post]
func CreateConsumer(c echo.Context) (*model.Consumer, error) {

	consumer := model.Consumer{}
	if error := c.Bind(&consumer); error != nil {
		return nil, error
	}
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
	result, err := arfcreateconsumer.ARF_CreateConsumer(&consumer, &cntxt, config)
	if reflect.TypeOf(result) == reflect.TypeOf(model.Consumer{}) || reflect.TypeOf(result) == reflect.TypeOf(&model.Consumer{}) {
		return result.(*model.Consumer), err
	} else {
		return nil, err
	}
}
