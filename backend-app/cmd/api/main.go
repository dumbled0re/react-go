package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	// flagでconfigのプロパティを初期化する
	// 引数は変数のポインタ(メモリのアドレス値)、フラグの名前、デフォルト値、使い方の説明
	flag.IntVar(&cfg.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment (development|production)")
	// Parse()でそれぞれの変数にアクセス可能
	flag.Parse()

	// Loggerオブジェクトを生成して出力フォーマットを設定する
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}

	// サーバー設定をカスタマイズ
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Println("Starting server on port", cfg.port)

	// :の前に何も書かなければローカルになる、第一引数にserverを立ち上げるポートを記載
	// 第二引数はハンドラーを渡す(nilの場合はnot-foundを返すハンドラーなのでその前にHandleFuncでハンドラーを渡す必要がある)
	err := srv.ListenAndServe()
	// err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.port), nil)
	if err != nil {
		log.Println(err)
	}
}
