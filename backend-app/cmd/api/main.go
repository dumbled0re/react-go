package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
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
	// flag.StringVar(&cfg.db.dsn, "dsn", "postgres://[role name]:[password]@localhost/[db name]?sslmode=disable", "Postgres connection string")
	flag.StringVar(&cfg.db.dsn, "dsn", "postgres://ritsushi:@localhost/go_movies?sslmode=disable", "Postgres connection string")
	// Parse()でそれぞれの変数にアクセス可能
	flag.Parse()

	// Loggerオブジェクトを生成して出力フォーマットを設定する
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

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
	err = srv.ListenAndServe()
	// err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.port), nil)
	if err != nil {
		log.Println(err)
	}
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	// コンテキスト&5秒のタイムアウト生成
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
