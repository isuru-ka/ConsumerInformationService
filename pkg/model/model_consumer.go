package model

import (
	"gorm.io/gorm"
)

type Consumer struct {
	Model                `bson:"-"`
	MID_                 string `bson:"_id,omitempty" gorm:"-" json:"-" xml:"mid_"`
	Consumername         string `bson:"consumername" json:"consumername" xml:"consumername"`
	Consumerid           string `bson:"consumerid" json:"consumerid" xml:"consumerid"`
	Clientname           string `bson:"clientname" json:"clientname" xml:"clientname"`
	Clientid             string `bson:"clientid" json:"clientid" xml:"clientid"`
	Cspid                string `bson:"cspid" json:"cspid" xml:"cspid"`
	Fedexid              string `bson:"fedexid" json:"fedexid" xml:"fedexid"`
	Clientrepresentative string `bson:"clientrepresentative" json:"clientrepresentative" xml:"clientrepresentative"`
	Ssn                  string `bson:"ssn" json:"ssn" xml:"ssn"`
	Othernames           string `bson:"othernames" json:"othernames" xml:"othernames"`
	Dateofbirth          string `bson:"dateofbirth" json:"dateofbirth" xml:"dateofbirth"`
	Address              string `bson:"address" json:"address" xml:"address"`
	Objectuuid           string `bson:"objectuuid" gorm:"objectuuid" json:"objectuuid" xml:"objectuuid"`
}

func (Consumer) TableName() string {
	return "consumer"
}
func (m *Consumer) PreloadConsumer(db *gorm.DB) *gorm.DB {
	return db
}
