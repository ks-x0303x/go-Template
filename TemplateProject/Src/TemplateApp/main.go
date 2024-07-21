package main

import (
	"Common/ExpansionString"
	"Common/Logger"
	"DatabaseManager"
	users "Models/users"
	"Server"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

// アプリケーションのエントリポイントです。
func main() {
	fmt.Println("Hello, World!")
	DatabaseManager.My_print()
	Logger.Instance.WriteLog("app start")

	// Create a wait group to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Launch multiple goroutines to write to the file concurrently
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			content := fmt.Sprintf("Hello from goroutine %d\n", i)
			Logger.Instance.WriteLog(content)
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	DatabaseManager.Instance.SetDBName("test_db")
	DatabaseManager.Instance.SetIpAddress("127.0.0.1")
	DatabaseManager.Instance.SetPort(13306)
	DatabaseManager.Instance.Connect()
	DatabaseManager.Instance.CanCommunication()

	//conditionInfo := DatabaseManager.NewQueryConditionInfo("ID", DatabaseManager.Equal, 2)
	conditionGroup := DatabaseManager.NewQueryConditionGroupInfo()
	//conditionGroup.Add(conditionInfo, DatabaseManager.None)
	recodes, err := users.Read(conditionGroup)
	if err != nil {
		fmt.Println("データなし")
	} else {
		for _, recode := range recodes {
			fmt.Println(ExpansionString.StructToString(recode))
		}

	}

	DatabaseManager.Instance.Disconnect()
	//db_test()
	server := Server.NewServer()
	stop := server.Start()
	Logger.Instance.WriteLog("server up")
	<-stop
	Logger.Instance.WriteLog("server stop")
	Logger.Instance.WriteLog("app end")
}
