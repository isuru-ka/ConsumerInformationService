package ARF_EditConsumer

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
	Consumer             *model.Consumer
	_xiDBNode_Password   interface{}
	consumerName         string
	cSPID                string
	fedexID              string
	clientRepresentative string
	sSN                  string
	otherNames           string
	_xiDBNode_UserName   interface{}
	address              string
	clientName           string
	dateOfBirth          string
	consumerUpdate       *model.Consumer
	consumerId           string
	_context             *util.CustomContext
	_ruleConfig          map[string]interface{}
	_returnValue         interface{}
	_errorValue          error
}

func NewCFGContext(pConsumer *model.Consumer, pContext *util.CustomContext, pRuleConfig map[string]interface{}) *CFGContext {
	return &CFGContext{

		Consumer:    pConsumer,
		_context:    pContext,
		_ruleConfig: pRuleConfig,
	}
}
func ARF_EditConsumer(pConsumer *model.Consumer, pContext *util.CustomContext, pRuleConfig map[string]interface{}) (interface{}, error) {

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

	cfg.xiReferencePropertyNodeM13()

	cfg.xiAssignNodeM34()

	if err :=
		cfg.xiDBNodeM42(); err != nil {

		return err
	}
	return nil
}

func (cfg *CFGContext) xiDBNodeM42() error {

	_, err := env.MDB.Collection("Consumer").UpdateOne(context.Background(), bson.M{"objectuuid": cfg.consumerUpdate.Objectuuid}, bson.D{{"$set", cfg.consumerUpdate}})
	if err != nil {

		xiLogger.Log.Error(err)
		cfg._errorValue = err
	}
	cfg._returnValue = cfg.consumerUpdate
	return nil
}

func (cfg *CFGContext) xiAssignNodeM34() error {

	cfg.consumerUpdate.Consumername = cfg.consumerName
	cfg.consumerUpdate.Clientname = cfg.clientName
	cfg.consumerUpdate.Cspid = cfg.cSPID
	cfg.consumerUpdate.Fedexid = cfg.fedexID
	cfg.consumerUpdate.Clientrepresentative = cfg.clientRepresentative
	cfg.consumerUpdate.Ssn = cfg.sSN
	cfg.consumerUpdate.Othernames = cfg.otherNames
	cfg.consumerUpdate.Dateofbirth = cfg.dateOfBirth
	cfg.consumerUpdate.Address = cfg.address

	cfg._returnValue = cfg.consumerUpdate.Address
	return nil
}

func (cfg *CFGContext) xiReferencePropertyNodeM13() error {

	return nil
}

func (cfg *CFGContext) xiDBNodeM01() error {

	_target := model.Consumer{}
	ctx := context.Background()
	opts := options.Find()
	cursor, err := env.MDB.Collection("Consumer").Find(context.Background(), bson.D{{"consumerid", cfg.consumerId}}, opts)
	if err != nil {

		xiLogger.Log.Error(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var data model.Consumer
		if err = cursor.Decode(&data); err != nil {
			xiLogger.Log.Error(err)
			break
		} else {
			_target = data
			break
		}
	}
	cfg.consumerUpdate = &_target
	cfg._returnValue = &_target
	return nil
}

func (cfg *CFGContext) xiModelPropertyNodeM0() error {
	cfg.consumerId = cfg.Consumer.Consumerid
	cfg.consumerName = cfg.Consumer.Consumername
	cfg.clientName = cfg.Consumer.Clientname
	cfg.cSPID = cfg.Consumer.Cspid
	cfg.fedexID = cfg.Consumer.Fedexid
	cfg.clientRepresentative = cfg.Consumer.Clientrepresentative
	cfg.sSN = cfg.Consumer.Ssn
	cfg.otherNames = cfg.Consumer.Othernames
	cfg.dateOfBirth = cfg.Consumer.Dateofbirth
	cfg.address = cfg.Consumer.Address
	cfg.consumerUpdate = &model.Consumer{}

	return nil
}
