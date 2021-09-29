package ARF_CreateConsumer

import util "consumerinformationmodule/pkg/util"

import (
	"consumerinformationmodule/pkg/env"
	"consumerinformationmodule/pkg/model"
	"consumerinformationmodule/pkg/xiLogger"
	"context"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CFGContext struct {
	Consumer           *model.Consumer
	consumerPropery    *model.Consumer
	_xiDBNode_UserName interface{}
	_xiDBNode_Password interface{}
	_context           *util.CustomContext
	_ruleConfig        map[string]interface{}
	_returnValue       interface{}
	_errorValue        error
}

func NewCFGContext(pConsumer *model.Consumer, pContext *util.CustomContext, pRuleConfig map[string]interface{}) *CFGContext {
	return &CFGContext{

		Consumer:    pConsumer,
		_context:    pContext,
		_ruleConfig: pRuleConfig,
	}
}
func ARF_CreateConsumer(pConsumer *model.Consumer, pContext *util.CustomContext, pRuleConfig map[string]interface{}) (interface{}, error) {

	cfg := NewCFGContext(pConsumer, pContext, pRuleConfig)
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

	objID := xid.New()
	cfg.consumerPropery.Objectuuid = objID.String()
	cfg.consumerPropery.MID_ = primitive.NewObjectID().Hex()
	_, err := env.MDB.Collection("Consumer").InsertOne(context.Background(), cfg.consumerPropery)
	if err != nil {

		xiLogger.Log.Error(err)
		cfg._errorValue = err
	}
	cfg._returnValue = cfg.consumerPropery
	return nil
}

func (cfg *CFGContext) xiModelPropertyNodeM0() error {
	cfg.consumerPropery = cfg.Consumer

	return nil
}
