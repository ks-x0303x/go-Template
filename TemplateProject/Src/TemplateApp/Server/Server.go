package Server

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/* #region struct */

type Human struct {
	Names []string
}

type Server struct {
	IPAddress  string
	Port       int
	httpServer *http.Server
}

/* #end region */

/* #region interface */

type ServerInterface interface {
	GetIpAddress() string
	SetIpAddress(ipAddress string)
	GetPort() int
	SetPort(port int)
	Start() chan os.Signal
	Stop() error
}

/* #end region */

/* #region コンストラクタ */

func NewServer() ServerInterface {
	return &Server{}
}

/* #end region */

/* #region Property */

// GetIpAddress method
func (server *Server) GetIpAddress() string {
	return server.IPAddress
}

// SetIpAddress method
func (server *Server) SetIpAddress(ipAddress string) {
	server.IPAddress = ipAddress
}

// GetPort method
func (server *Server) GetPort() int {
	return server.Port
}

// SetPort method
func (server *Server) SetPort(port int) {
	server.Port = port
}

/* #end region */

/* #region Method */

/*	-------------------------------------------- */
// 概要	: サーバーをスタートします。
// 引数	: -
// 戻値	: osのSignal interface
// 備考	:
/*	-------------------------------------------- */
func (server *Server) Start() chan os.Signal {
	serveMux := http.NewServeMux()
	// ルート名　コールバック関数登録
	serveMux.HandleFunc("/view", viewHandler)

	// シグナルをキャッチするためのチャネルを作成します
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// シャットダウンエンドポイントを登録
	serveMux.HandleFunc("/shutdown", shutdownHandler(stop))

	server.httpServer = &http.Server{
		Addr:    "0.0.0.0:8080", // 全てのネットワークインターフェースでリッスン
		Handler: serveMux,
	}

	// サーバースタート
	// サーバーを別のゴルーチンで実行します
	go func() {
		if err := server.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	return stop
}

func (server *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return server.httpServer.Shutdown(ctx)
}

/* #end region */

/* #region Func */

/*	-------------------------------------------- */
// 概要	: view のリクエストを作成します。
// 引数	: http.ResponseWriter
// 	　　: - HTTP ハンドラーによって HTTP レスポンスを作成するIF
// 	　　: http.Request
// 	　　: - サーバーが受信した HTTP リクエスト、またはクライアントが送信する HTTP リクエストを表す。
// 戻値	: -
// 備考	:
/*	-------------------------------------------- */
func viewHandler(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("./Server/view.html")
	if err != nil {
		log.Fatal(err)
	}

	human := Human{Names: []string{"seiya"}}

	if err := html.Execute(w, human); err != nil {
		log.Fatal(err)
	}

}

/*	-------------------------------------------- */
// 概要	: shutdownHandler リクエストを作成します。
// 引数	: stop
// 	　　: - 止めるときのチャネル
// 戻値	: コルバック関数
// 備考	:
/*	-------------------------------------------- */
func shutdownHandler(stop chan os.Signal) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		go func() {
			log.Println("Shutdown request received")
			stop <- syscall.SIGTERM
		}()
		w.Write([]byte("Server is shutting down"))
	}
}

// func CreateHandler(w http.ResponseWriter, r *http.Request) {
// 	formValue := r.FormValue("value")
// 	flag := os.O_WRONLY | os.O_APPEND | os.O_CREATE
// 	// fileMode unix系のファイル権限を表す　右から　特集：所有者：グループ：ゲスト(全員)
// 	// 4：読込権限　２：書込権限　１：実行権限
// 	file, err := os.OpenFile(fileName, flag, os.FileMode(0600))
// 	defer file.Close()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	_, err = fmt.Fprintln(file, formValue)
// 	if err != nil {
// 		log.Fatal()
// 	}
// 	http.Redirect(w, r, "/view", http.StatusFound)
// }

/* #end region */
