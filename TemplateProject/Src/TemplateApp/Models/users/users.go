package users

import (
	"Common/Logger"
	"DatabaseManager"
	"fmt"
	"reflect"

	"time"
)

type users struct {
	ID        int
	FirstName string
	LastName  string
	Age       string
	Created   time.Time
	Updated   time.Time
}

func Create(value interface{}) error {
	return nil
}

func Read(conditionGroupInfo DatabaseManager.QueryConditionGroupInfoInterface) ([]interface{}, error) {
	elementType := reflect.TypeOf(users{})
	query := conditionGroupInfo.CreateReadQuery(elementType.Name())
	fmt.Println(query)
	return DatabaseManager.Instance.GetRecords(elementType, query)
}

func Update(value interface{}) error {
	Logger.Instance.WriteLog("")
	return nil
}

func Delete(conditionGroupInfo *DatabaseManager.QueryConditionGroupInfoInterface) error {
	return nil
}
