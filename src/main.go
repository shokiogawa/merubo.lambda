package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log/slog"
)

// 参考
// https://yhidetoshi.hatenablog.com/entry/2022/10/10/145121

//var echoLambda *echoadapter.EchoLambda

func init() {
	slog.Info("コールドスタート")
	//e := echo.New()
	//e.Use(middleware.Recover())
	//e.GET("/api/contact", GetContact)
	//e.POST("/api/contact", CreateContact)
	//
	//// セッティング
	//echoLambda = echoadapter.New(e)
}

func main() {
	slog.Info("Goアプリ開始")
	lambda.Start(HandleRequest)
}

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	slog.Info("HandlerRequest")
	return events.APIGatewayProxyResponse{
		Body:       "Hello Go",
		Headers:    nil,
		StatusCode: 200}, nil
}

func TestFunc() {

}
