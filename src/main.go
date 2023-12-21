package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log/slog"
	"net/http"
)

// 参考
// https://yhidetoshi.hatenablog.com/entry/2022/10/10/145121

var echoLambda *echoadapter.EchoLambda

func init() {
	slog.Info("コールドスタート")
	e := echo.New()
	e.Use(middleware.Recover())
	e.POST("/api/contact", PostContact)
	e.GET("/api/contact", GetContact)

	echoLambda = echoadapter.New(e)
}

func PostContact(e echo.Context) (err error) {
	slog.Info("PostContact")
	return e.String(http.StatusCreated, "PostContact success")
}

func GetContact(e echo.Context) (err error) {
	slog.Info("GetContact")
	return e.String(http.StatusOK, "GetContent success")
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	slog.Info(req.Path)
	slog.Info(req.HTTPMethod)
	return echoLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
	//init, err := NewInitialize()
	//if err != nil {
	//	slog.Error(err.Error(), "message", "初期化失敗")
	//}
	//if os.Getenv("ENV") == "dev" {
	//	slog.Info("開発環境")
	//	e := echo.New()
	//	//e.POST("/callback", init.LambdaHandler.LineHandler.EventHandler)
	//	e.POST("/api/contact", init.LambdaHandler.LineHandler.SendContactMessageEcho)
	//	e.Logger.Fatal(e.Start(":8080"))
	//} else {
	//	slog.Info("本番環境")
	//	lambda.Start(init.LambdaHandler.HandleRequest)
	//}
}
