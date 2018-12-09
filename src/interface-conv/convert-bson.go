package interface_conv

import (
	"github.com/globalsign/mgo/bson"
	"interfaces-internal"
)

func ConvertBSONToIAccount(original bson.M) (interfaces_internal.IAccount, error) {

	account := interfaces_internal.IAccount{}
	bsonBytes, err := bson.Marshal(original)
	if err != nil {
		return account, err
	}
	err = bson.Unmarshal(bsonBytes, &account)
	if err != nil {
		return account, err
	}

	return account, nil
}

func ConvertBSONToIOrganization(original bson.M) (interfaces_internal.IOrganization, error) {

	account := interfaces_internal.IOrganization{}
	bsonBytes, err := bson.Marshal(original)
	if err != nil {
		return account, err
	}
	err = bson.Unmarshal(bsonBytes, &account)
	if err != nil {
		return account, err
	}

	return account, nil
}