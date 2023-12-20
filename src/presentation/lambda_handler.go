package presentation

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"merubo.lambda/src/presentation/client_model/request_model"
)

type LambdaHandler struct {
	LineHandler *LineHandler
}

func NewLambdaHandler(lineHandler *LineHandler) (handler *LambdaHandler) {
	handler = new(LambdaHandler)
	handler.LineHandler = lineHandler
	return
}

func (lambdaHandler *LambdaHandler) HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (response events.APIGatewayProxyResponse, err error) {
	contactBody, err := request_model.ConvertContactRequestModel(request.Body)
	if err != nil {
		return
	}
	err = lambdaHandler.LineHandler.SendContactMessageLambda(contactBody)
	if err != nil {
		return
	}
	return events.APIGatewayProxyResponse{
		Body:       "問い合わせ成功",
		Headers:    nil,
		StatusCode: 200}, nil
}
