package ARF_DeleteConsumers

import util "consumerinformationmodule/pkg/util"

import (
	"consumerinformationmodule/pkg/env"
	"consumerinformationmodule/pkg/xiLogger"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type CFGContext struct {
	_xiDBNode_UserName interface{}
	_xiDBNode_Password interface{}
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
func ARF_DeleteConsumers(pContext *util.CustomContext, pRuleConfig map[string]interface{}) (interface{}, error) {

	cfg := NewCFGContext(pContext, pRuleConfig)
	cfg.r0()
	return cfg._returnValue, cfg._errorValue
}
func (cfg *CFGContext) r0() error {

	if err :=
		cfg.xiDBNodeM0(); err != nil {

		return err
	}
	return nil
}

func (cfg *CFGContext) xiDBNodeM0() error {

	_, err := env.MDB.Collection("Consumer").DeleteMany(context.Background(), bson.M{})
	if err != nil {

		xiLogger.Log.Error(err)
	}
	return nil
}
