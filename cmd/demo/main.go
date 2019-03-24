package main

import (
	"fmt"
	"runtime"

	log "github.com/sirupsen/logrus"
	logger "github.com/zylisp/zylog/logger"
)

func SetupLogger() {
	logger.SetupLogging(&logger.ZyLogOptions{
		Colored:      true,
		Level:        "trace",
		Output:       "stdout",
		ReportCaller: true,
	})
}

func SetupLoggerNoCaller() {
	logger.SetupLogging(&logger.ZyLogOptions{
		Colored:      true,
		Level:        "trace",
		Output:       "stdout",
		ReportCaller: false,
	})
}

func printVersions() {
	fmt.Printf("zylog version: %s\n", logger.VersionString())
	fmt.Printf("Build: %s\n", logger.BuildString())
	fmt.Printf("Go version: %s\n", runtime.Version())
}

func main() {
	printVersions()
	SetupLogger()
	log.Trace("This is trace")
	log.Debug("This is debug")
	log.Info("This is info")
	log.Warn("This is warn")
	log.Error("This is error")
	log.Info("Fatal and Panic are also supported; "+
		"Fatal will os.Exit, and Panic will log, then panic().")
	log.Info("When not testing, you'll want to turn off caller reporting:")
	SetupLoggerNoCaller()
	log.Trace("This is trace")
	log.Debug("This is debug")
	log.Info("This is info")
	log.Warn("This is warn")
	log.Error("This is error")
}
