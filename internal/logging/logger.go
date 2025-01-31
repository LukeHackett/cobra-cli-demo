package logging

import (
	"log"
	"log/slog"
	"math"
)

func ConfigureLogging(debug bool) {
	if debug {
		log.SetFlags(log.LstdFlags)
		slog.SetLogLoggerLevel(slog.LevelDebug)
	} else {
		log.SetFlags(0) // turn off all additional flags
		slog.SetLogLoggerLevel(math.MaxInt)
	}
}
