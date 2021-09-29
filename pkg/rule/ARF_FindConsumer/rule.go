package ARF_FindConsumer

import util "consumerinformationmodule/pkg/util"

import (
	"consumerinformationmodule/pkg/env"
	"consumerinformationmodule/pkg/model"
	"consumerinformationmodule/pkg/xiLogger"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CFGContext struct {
	consumerid         string
	_xiDBNode_Password interface{}
	consumerProperty   []*model.Consumer
	consumerList       []*model.Consumer
	_xiDBNode_UserName interface{}
	_context           *util.CustomContext
	_ruleConfig        map[string]interface{}
	_returnValue       interface{}
	_errorValue        error
}

func NewCFGContext(pConsumerid string, pContext *util.CustomContext, pRuleConfig map[string]interface{}) *CFGContext {
	return &CFGContext{

		consumerid:  pConsumerid,
		_context:    pContext,
		_ruleConfig: pRuleConfig,
	}
}
func ARF_FindConsumer(pConsumerid string, pContext *util.CustomContext, pRuleConfig map[string]interface{}) (interface{}, error) {

	cfg := NewCFGContext(pConsumerid, pContext, pRuleConfig)
	cfg.r0()
	return cfg._returnValue, cfg._errorValue
}
func (cfg *CFGContext) r0() error {

	cfg.xiModelPropertyNodeM0()

	if err :=
		cfg.xiDBNodeM10(); err != nil {

		return err
	}
	return nil
}

func (cfg *CFGContext) xiDBNodeM10() error {

	_target := []*model.Consumer{}
	ctx := context.Background()
	opts := options.Find()
	cursor, err := env.MDB.Collection("Consumer").Find(context.Background(), bson.D{{"consumerid", cfg.consumerid}}, opts)
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
	cfg.consumerList = _target
	cfg._returnValue = _target
	return nil
}

func (cfg *CFGContext) xiModelPropertyNodeM0() error {

	return nil
}
