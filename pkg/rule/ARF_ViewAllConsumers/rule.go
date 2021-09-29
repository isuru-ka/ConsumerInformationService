package ARF_ViewAllConsumers

import util "consumerinformationmodule/pkg/util"

import (
	"consumerinformationmodule/pkg/env"
	"consumerinformationmodule/pkg/model"
	"consumerinformationmodule/pkg/xiLogger"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type CFGContext struct {
	_xiDBNode_UserName interface{}
	_xiDBNode_Password interface{}
	cp                 []*model.Consumer
	_context           *util.CustomContext
	_ruleConfig        map[string]interface{}
	_returnValue       interface{}
	_errorValue        error
}

func NewCFGContext(pContext *util.CustomContext, pRuleConfig map[string]interface{}) *CFGContext {
	return &CFGContext{

		_context:    pContext,
		_ruleConfig: pRuleConfig,
	}
}
func ARF_ViewAllConsumers(pContext *util.CustomContext, pRuleConfig map[string]interface{}) (interface{}, error) {

	cfg := NewCFGContext(pContext, pRuleConfig)
	cfg.r0()
	return cfg._returnValue, cfg._errorValue
}
func (cfg *CFGContext) r0() error {

	cfg.xiModelPropertyNodeM0()

	if err :=
		cfg.xiDBNodeM01(); err != nil {

		return err
	}
	return nil
}

func (cfg *CFGContext) xiDBNodeM01() error {

	_target := []*model.Consumer{}
	ctx := context.Background()
	cursor, err := env.MDB.Collection("Consumer").Find(context.Background(), bson.M{})
	if err != nil {

		xiLogger.Log.Error(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var data model.Consumer
		if err = cursor.Decode(&data); err != nil {
			xiLogger.Log.Error(err)
		} else {
			_target = append(_target, &data)
		}
	}
	cfg.cp = _target
	cfg._returnValue = _target
	return nil
}

func (cfg *CFGContext) xiModelPropertyNodeM0() error {

	return nil
}
