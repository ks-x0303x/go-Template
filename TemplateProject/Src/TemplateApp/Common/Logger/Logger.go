package Logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var (
	// Log Instance
	Instance LoggerInterface = new(LoggerInstance)

	// 排他制御オブジェクト
	mu sync.Mutex
)

// Log Instance
type LoggerInstance struct {
	FileName     string
	FolderPath   string
	IsInitialize bool
}

// Log interface
type LoggerInterface interface {
	WriteLog(text string) bool
	TraceLog(err error)
}

const logDir = "log"
const extension = "." + logDir

// 初期化します
func (logger *LoggerInstance) Initialize() bool {
	// ディレクトリ確認
	if f, err := os.Stat("./" + logDir); os.IsNotExist(err) || !f.IsDir() {
		// ディレクトリ作成
		fileInfo, err := os.Lstat("./")
		if err != nil {
			return logger.IsInitialize
		}
		fileMode := fileInfo.Mode()
		unixPerms := fileMode & os.ModePerm
		if err := os.MkdirAll(logDir, unixPerms); err != nil {
			return logger.IsInitialize
		}
	}

	logger.IsInitialize = true
	return logger.IsInitialize
}

// Logを書き込みます
func (logger *LoggerInstance) WriteLog(text string) bool {
	// 本関数のクリティカルセクションをロック (ファイルアクセス)
	// 書き込みロック
	mu.Lock()
	defer mu.Unlock()
	if !logger.IsInitialize {
		//return false
		if !logger.Initialize() {
			return logger.IsInitialize
		}
	}
	var fileName string = ""
	if logger.FileName == "" {
		fileName = time.Now().Format("20060102") + extension
	} else {
		fileName = logger.FileName + extension
	}
	// ファイル開く
	file, err := os.OpenFile(logDir+"/"+fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	// 関数終了時またはエラー発生時ファイル閉じる
	defer file.Close()

	// ログをファイル出力する
	log.SetOutput(file)
	// ログ出力
	var pid string = strconv.Itoa(os.Getpid())
	log.Print("PID=" + pid + " " + text)
	return true
}

func (logger *LoggerInstance) TraceLog(err error) {

	// スタックトレースの情報を取得
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("エラー情報の取得に失敗しました")
		return
	}

	// 関数名を取得
	fn := runtime.FuncForPC(pc)
	funcName := "unknown"
	if fn != nil {
		funcName = fn.Name()
	}

	// エラーメッセージを表示
	fmt.Printf("エラー: %s\nファイル: %s\n関数: %s\n行: %d\n", err, file, funcName, line)
}
