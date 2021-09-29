package functions

import (
	"consumerinformationmodule/pkg/model"
)

func FilterCompany(consumerlist []*model.Consumer) (*model.Consumer, error) {
 if len(consumerlist) > 0 {

        consumerOut := consumerlist[0]
        return consumerOut, nil

    } else {

        return nil, nil
    }
}
