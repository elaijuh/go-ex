package main

import (
	"flag"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

func main() {
	debug := flag.Bool("debug", false, "debug level")
	flag.Parse()
	// logger := log.NewLogfmtLogger(os.Stdout)
	logger := log.NewJSONLogger(os.Stdout)

	// logger = log.With(logger, "time", log.DefaultTimestampUTC, "caller", log.DefaultCaller)
	logger = log.With(logger, "time", log.DefaultTimestampUTC, "caller", log.Caller(5))

	_ = debug

	if *debug {
		logger = level.NewFilter(logger, level.AllowDebug())
	} else {
		logger = level.NewFilter(logger, level.AllowInfo())
	}

	level.Info(logger).Log("foo", "bar", "event", "test")
	level.Debug(logger).Log("debug", "true")

	loggerp := &logger
	level.Info(*loggerp).Log("pointer", "pointer")
	level.Debug(*loggerp).Log("pointer", "pointer")

}
