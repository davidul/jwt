package main

import (
	"go.uber.org/zap"
	"jwt/cmd"
)

func main() {
	logger := zap.NewExample()
	defer logger.Sync()

	err := cmd.Execute(logger)
	if err != nil {
		return
	}
}
