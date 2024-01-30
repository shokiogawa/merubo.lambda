package lib

import (
	"log/slog"
	"os"
)

func NewLogger() (logger *slog.Logger) {
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	return logger
}

// testb
func testb(){}

// テストA
func testA(){}