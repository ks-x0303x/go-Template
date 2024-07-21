package DatabaseManager

import (
	"Common/Logger"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
)

var (
	Instance = NewSMySqLManager()
)

/* #region struct */

type MySqLManager struct {
	DBName    string
	IPAddress string
	Port      int
	IsConnect bool
	MySqL     *sql.DB
}

/* #end region */

/* #region interface */

type MySqLManagerInterface interface {
	// DB name
	GetDBName() string
	SetDBName(name string)

	// Ip address
	GetIpAddress() string
	SetIpAddress(ipAddress string)

	// Port
	GetPort() int
	SetPort(port int)

	// DB接続状態
	GetIsConnect() bool

	Connect() bool
	CanCommunication() bool
	Disconnect()
	GetRecords(elementType reflect.Type, query string) ([]interface{}, error)
}

/* #end region */

/* #region コンストラクタ */

func NewSMySqLManager() MySqLManagerInterface {
	return &MySqLManager{}
}

/* #end region */

/* #region Property */

// GetDBName method
func (mysqlManager *MySqLManager) GetDBName() string {
	return mysqlManager.DBName
}

// SetIDBName method
func (mysqlManager *MySqLManager) SetDBName(name string) {
	mysqlManager.DBName = name
}

// GetIpAddress method
func (mysqlManager *MySqLManager) GetIpAddress() string {
	return mysqlManager.IPAddress
}

// SetIpAddress method
func (mysqlManager *MySqLManager) SetIpAddress(ipAddress string) {
	mysqlManager.IPAddress = ipAddress
}

// GetPort method
func (mysqlManager *MySqLManager) GetPort() int {
	return mysqlManager.Port
}

// SetPort method
func (mysqlManager *MySqLManager) SetPort(port int) {
	mysqlManager.Port = port
}

// IsConnect method
func (mysqlManager *MySqLManager) GetIsConnect() bool {
	return mysqlManager.IsConnect
}

/* #end region */

/* #region Method */

/*	-------------------------------------------- */
// 概要	: DBに接続します。
// 引数	: -
// 戻値	: true 成功　：　false 失敗
// 備考	:
/*	-------------------------------------------- */
func (mysqlManager *MySqLManager) Connect() bool {
	mysqlManager.IsConnect = false

	// Check ip address
	var ipAddress = mysqlManager.IPAddress
	if ipAddress == "" {
		Logger.Instance.WriteLog("IpAddress is not set.")
		return mysqlManager.IsConnect
	}

	// Check port number
	var portNum = mysqlManager.Port
	var port = strconv.Itoa(portNum)
	if portNum == 0 {
		Logger.Instance.WriteLog("Port is not set.")
		return mysqlManager.IsConnect
	}

	// db connect
	var address = ipAddress + ":" + port
	db, err := sql.Open("mysql", "test_user:test_password@tcp("+address+")/test_db?parseTime=true&loc=Asia%2FTokyo")
	mysqlManager.MySqL = db
	if err != nil {
		Logger.Instance.WriteLog(" MySqL is Connect error. : " + err.Error())
		return mysqlManager.IsConnect
	} else {
		mysqlManager.IsConnect = true
	}
	Logger.Instance.WriteLog("DB接続 : " + address)
	return mysqlManager.IsConnect
}

/*	-------------------------------------------- */
// 概要	: DBと通信が取れるか確認します。
// 引数	: -
// 戻値	: true 通信可能　：　false 通信不可
// 備考	:
/*	-------------------------------------------- */
func (mysqlManager *MySqLManager) CanCommunication() bool {

	if !mysqlManager.IsConnect {
		Logger.Instance.WriteLog("DB接続されていません。")
		return false
	}
	err := mysqlManager.MySqL.Ping()
	if err != nil {
		Logger.Instance.WriteLog("DBのping応答なし")
		return false
	}
	Logger.Instance.WriteLog("DBのping応答あり")
	return true
}

/*	-------------------------------------------- */
// 概要	: DBを切断します。
// 引数	: -
// 戻値	: -
// 備考	:
/*	-------------------------------------------- */
func (mysqlManager *MySqLManager) Disconnect() {
	mysqlManager.MySqL.Close()
	Logger.Instance.WriteLog("DB切断")
}

/*	-------------------------------------------- */
// 概要	: レコードを取得します。
// 引数	: -
// 戻値	: -
// 備考	:
/*	-------------------------------------------- */
func (mysqlManager *MySqLManager) GetRecords(elementType reflect.Type, query string) ([]interface{}, error) {
	// クエリを実行
	rows, err := mysqlManager.MySqL.Query(query)
	if err != nil {
		Logger.Instance.WriteLog("getRecords db.Query error: " + err.Error())
		return nil, err
	}
	defer rows.Close()

	// 列のスキャン
	var records []interface{}
	for rows.Next() {
		// 新しい構造体を作成
		elem := reflect.New(elementType).Elem()
		fields := make([]interface{}, elem.NumField())
		for i := 0; i < elem.NumField(); i++ {
			fields[i] = elem.Field(i).Addr().Interface()
		}

		if err := rows.Scan(fields...); err != nil {
			Logger.Instance.WriteLog("getRecords rows.Scan error: " + err.Error())
			return nil, err
		}

		// レコードをスライスに追加
		records = append(records, elem.Interface())
	}

	if err := rows.Err(); err != nil {
		Logger.Instance.WriteLog("getRecords rows.Err error: " + err.Error())
		return nil, err
	}

	return records, nil
}

func My_print() {
	fmt.Println("Databasemanager")
}

/* #end region */
