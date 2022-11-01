package bootstrap

import "github.com/feacx/study/pkg/logger"

func SetupLogger() {
	logger.InitLogger("storage/logs/app.log")
}
