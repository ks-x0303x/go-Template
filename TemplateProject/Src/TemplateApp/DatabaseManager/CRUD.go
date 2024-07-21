package DatabaseManager

type CRUD_interface interface {
	Create(value interface{}) error
	Read(conditionGroupInfo *QueryConditionGroupInfoInterface) ([]interface{}, error)
	Update(value interface{}) error
	Delete(conditionGroupInfo *QueryConditionGroupInfoInterface) error
}
