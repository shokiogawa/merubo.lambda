package main

import (
	"merubo.lambda/src/lib"
	"merubo.lambda/src/presentation"
)

type Initialize struct {
	LambdaHandler *presentation.LambdaHandler
}

func NewInitialize() (init *Initialize, err error) {
	init = new(Initialize)
	logger := lib.NewLogger()
	lineBot, err := lib.NewLineBot()
	lineHandler := presentation.NewLineHandler(logger, lineBot)

	lambdaHandler := presentation.NewLambdaHandler(lineHandler)
	init.LambdaHandler = lambdaHandler
	return
}
