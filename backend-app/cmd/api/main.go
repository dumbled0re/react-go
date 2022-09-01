package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
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

func main() {
	var cfg config

	// 型名Var()を使った場合、引数で渡した変数にバインド
	flag.IntVar(&cfg.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment (development|production)")
	// Parse()でそれぞれの変数にアクセス可能
	flag.Parse()

	fmt.Println("Running")

	// HandleFuncはルーティング先とHandler関数を渡す必要あり
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		currentStatus := AppStatus{
			Status:      "Available",
			Environment: cfg.env,
			Version:     version,
		}

		// 構造体をjsonに変換
		js, err := json.MarshalIndent(currentStatus, "", "\t")
		if err != nil {
			log.Println(err)
		}

		// この順番ではないとエラー発生する
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(js)
	})

	// :の前に何も書かなければローカルになる、第一引数にserverを立ち上げるポートを記載
	// 第二引数はハンドラーを渡す(nilの場合はnot-foundを返すハンドラーなのでその前にHandleFuncでハンドラーを渡す必要がある)
	err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.port), nil)
	if err != nil {
		log.Println(err)
	}
}
