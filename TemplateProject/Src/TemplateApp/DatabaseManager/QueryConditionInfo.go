package DatabaseManager

import (
	"Common/Logger"
	"errors"
	"fmt"
)

type ConditionType string
type CompoundConditionType string

const (
	Equal              ConditionType = "="
	LessThan           ConditionType = "<"
	GreaterThan        ConditionType = ">"
	NotEqual           ConditionType = "!="
	LessThanOrEqual    ConditionType = "<="
	GreaterThanOrEqual ConditionType = ">="
)

const (
	None CompoundConditionType = ""
	And  CompoundConditionType = "AND"
	Or   CompoundConditionType = "OR"
)

/* #region struct */

type QueryConditionInfo struct {
	FieldName         string
	Condition         ConditionType
	Value             string
	CompoundCondition CompoundConditionType
}

type QueryConditionGroupInfo struct {
	ConditionInfoList []*QueryConditionInfo
}

/* #end region */

/* #region interface */

type QueryConditionInfoInterface interface {
}

type QueryConditionGroupInfoInterface interface {
	Add(conditionInfo *QueryConditionInfo, compoundConditionType CompoundConditionType) error
	GetConditionInfoList() []*QueryConditionInfo
	CreateReadQuery(tableName string) string
}

/* #end region */

/* #region コンストラクタ */

func NewQueryConditionInfo(fieldName string, condition ConditionType, value interface{}) *QueryConditionInfo {
	return &QueryConditionInfo{
		FieldName: fieldName,
		Condition: condition,
		Value:     GetValue(value),
	}
}

func NewQueryConditionGroupInfo() QueryConditionGroupInfoInterface {
	return new(QueryConditionGroupInfo)
}

/* #end region */

/* #region QueryConditionInfo Property */

// Set CompoundCondition method
func (conditionInfo *QueryConditionInfo) SetCompoundCondition(condition CompoundConditionType) {
	conditionInfo.CompoundCondition = condition
}

// Get CompoundCondition method
func (conditionInfo *QueryConditionInfo) GetCompoundCondition() CompoundConditionType {
	return conditionInfo.CompoundCondition
}

/* #end region */

/* #region QueryConditionInfo Method */
/* #end region */

/* #region QueryConditionGroupInfo Property */
/* #end region */

/* #region QueryConditionGroupInfo Method */

/*	-------------------------------------------- */
// 概要	: クエリに含む条件を追加します
// 引数	: conditionInfo
//		:  - クエリ条件
//		: compoundConditionType
//		:  - 複合条件
// 戻値	: err
// 備考	:
/*	-------------------------------------------- */
func (conditionGroupInfo *QueryConditionGroupInfo) Add(conditionInfo *QueryConditionInfo, compoundConditionType CompoundConditionType) error {
	// nilチェック
	if conditionInfo == nil {
		err := errors.New("conditionInfo is nil")
		Logger.Instance.TraceLog(err)
		return err
	}
	// ２つ目以降の時に複合条件をNONEにすることはできない
	if (len(conditionGroupInfo.ConditionInfoList) > 0) && compoundConditionType == None {
		err := errors.New("2個以上の条件を使用する場合は、複合条件をANDかORを指定してください。")
		Logger.Instance.TraceLog(err)
		return err
	}
	// クエリグループにクエリ追加
	conditionInfo.CompoundCondition = compoundConditionType
	conditionGroupInfo.ConditionInfoList = append(conditionGroupInfo.ConditionInfoList, conditionInfo)
	return nil
}

/*	-------------------------------------------- */
// 概要	: クエリグループを取得します。
// 引数	: -
// 戻値	: クエリグループ
// 備考	:
/*	-------------------------------------------- */
func (conditionGroupInfo *QueryConditionGroupInfo) GetConditionInfoList() []*QueryConditionInfo {
	return conditionGroupInfo.ConditionInfoList
}

/*	-------------------------------------------- */
// 概要	: クエリを取得します
// 引数	: クエリグループ
// 戻値	: クエリ
// 備考	:
/*	-------------------------------------------- */
func (conditionGroupInfo *QueryConditionGroupInfo) CreateReadQuery(tableName string) string {

	var result = "SELECT * FROM " + Instance.GetDBName() + "." + tableName
	var condition = ""

	for _, info := range conditionGroupInfo.ConditionInfoList {

		// 複合条件
		condition += " " + string(info.CompoundCondition)
		condition += " ( "
		condition += info.FieldName
		condition += string(info.Condition)
		condition += info.Value
		condition += " ) "
	}
	if len(conditionGroupInfo.ConditionInfoList) > 0 {
		result += " WHERE " + condition
	}
	return result
}

/* #end region */

/* #region Function */

func GetValue(value interface{}) string {

	switch data := value.(type) {
	case string:
		return fmt.Sprintf("'%s'", data)
	default:
		return fmt.Sprintf("%v", data)
	}
}

/* #end region */
