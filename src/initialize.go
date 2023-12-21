package main

import (
	"merubo.lambda/src/lib"
	"merubo.lambda/src/presentation"
)

type Initialize struct {
	LineHandler *presentation.LineHandler
}

func NewInitialize() (init *Initialize, err error) {
	init = new(Initialize)
	logger := lib.NewLogger()
	lineBot, err := lib.NewLineBot()
	init.LineHandler = presentation.NewLineHandler(logger, lineBot)
	return
}
