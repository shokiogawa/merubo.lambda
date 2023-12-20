package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/labstack/echo/v4"
	"log/slog"
	"os"
)

// 参考
// https://yhidetoshi.hatenablog.com/entry/2022/10/10/145121

//var echoLambda *echoadapter.EchoLambda

func init() {
	slog.Info("コールドスタート")
}

func main() {
	init, err := NewInitialize()
	if err != nil {
		slog.Error(err.Error(), "message", "初期化失敗")
	}
	if os.Getenv("ENV") == "dev" {
		slog.Info("開発環境")
		e := echo.New()
		//e.POST("/callback", init.LambdaHandler.LineHandler.EventHandler)
		e.POST("/api/contact", init.LambdaHandler.LineHandler.SendContactMessageEcho)
		e.Logger.Fatal(e.Start(":8080"))
	} else {
		slog.Info("本番環境")
		lambda.Start(init.LambdaHandler.HandleRequest)
	}
}
