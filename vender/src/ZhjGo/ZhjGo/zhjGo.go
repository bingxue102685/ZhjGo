package ZhjGo

import (
	"ZhjGo/ZhjGo/conf"
	"ZhjGo/ZhjGo/log"
)

func Run() {
	if conf.LogLevel != "" {
		logger, err := log.New(conf.LogLevel, conf.LogPath, conf.LogFlag)
		if err != nil {
			panic(err)
		}
		log.Export(logger)
		defer logger.Close()
	}
	log.Release("ZhjGo %v starting up", version)
}
