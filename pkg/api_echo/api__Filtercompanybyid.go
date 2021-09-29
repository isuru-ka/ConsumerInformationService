package api_echo

import (
	model "consumerinformationmodule/pkg/model"
	arffiltercompanybyid "consumerinformationmodule/pkg/rule/ARF_FilterCompanyById"
	util "consumerinformationmodule/pkg/util"
	"context"
	"reflect"
	"strconv"
)

import (
	"github.com/labstack/echo/v4"
)

// @Summary FilterCompanyById
// @Tags consumerinformationmodule
// @Accept json
// @Produce json
// @Param consumerid query string true  " string consumerid query "
// @Success 200 {object} controller_echo.ListConsumerResponseFilterCompanyById
// @Router /consumerinformationmodule/api/consumer/profile [get]
func FilterCompanyById(c echo.Context) ([]*model.Consumer, error) {

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
	result, err := arffiltercompanybyid.ARF_FilterCompanyById(consumerid, &cntxt, config)
	if reflect.TypeOf(result) == reflect.TypeOf([]*model.Consumer{}) || reflect.TypeOf(result) == reflect.TypeOf([]*model.Consumer{}) {
		return result.([]*model.Consumer), err
	} else {
		return nil, err
	}
}
