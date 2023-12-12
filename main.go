package main

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"github.com/yizeng/toggl-test-signer/internal/config"
	"github.com/yizeng/toggl-test-signer/internal/logger"
	"github.com/yizeng/toggl-test-signer/internal/web"
)

func main() {
	err := config.Init()
	if err != nil {
		panic(fmt.Sprintf("failed to initialize config: %v", err))
	}

	err = logger.Init()
	if err != nil {
		panic(fmt.Sprintf("failed to initialize logger: %v", err))
	}

	s := web.NewServer()

	zap.L().Info("starting server at...", zap.String("address", s.Address))
	err = http.ListenAndServe(s.Address, s.Router)
	if err != nil {
		panic(fmt.Sprintf("failed to start the server: %v", err))
	}
}
