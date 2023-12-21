package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/labstack/echo/v4/middleware"
	"log/slog"
	"os"
)

var echoLambda *echoadapter.EchoLambda

func init() {
	slog.Info("コールドスタート")
	//依存関係解消
	init, err := NewInitialize()
	if err != nil {

	}
	// ルーター設定
	e := NewRouter(init)
	e.Use(middleware.Recover())

	// 開発環境の場合は、8080ポートにアクセス
	if os.Getenv("ENV") == "dev" {
		e.Logger.Fatal(e.Start(":8080"))
	}
	echoLambda = echoadapter.New(e)
}

func main() {
	lambda.Start(Handler)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return echoLambda.ProxyWithContext(ctx, req)
}
