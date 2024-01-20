package logcfg

import (
	"os"

	log "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

func Setup(level string, path string) {
	parsedLevel, err := log.ParseLevel(level)
	if err != nil {
		log.Fatalf("Invalid log level: %s", level)
	}

	log.SetLevel(parsedLevel)

	if path != "" {
		file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("Error opening log file: %v", err)
		}

		log.SetOutput(file)
	} else {
		log.SetOutput(os.Stdout)
	}

	log.SetFormatter(&easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[%lvl%]: %time% | %msg%\n",
	})
}
