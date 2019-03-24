/*
This package performs basic setup of the logrus library with custom formatting.

Overview

Zylog logger's primary features include:

	* Exceedingly simple setup
	* Colored output (enabled/disabled with a boolean)
	* Logging level (lower-case string)
	* Output (only stdout and stderr currently supported)
	* ReportCaller (enabled/disabled with a boolean; prints package, function
	  and line number)
	* Custom format (similar to the Clojure twig library and the LFE logjam
		libraries)

Setup is done with the zylog logger, after which logrus may be used as designed
by its author.

Installation

	$ go get github.com/zylisp/zylog/logger

Additionally, there is a demo you may install and run:

	$ go get github.com/zylisp/zylog/cmd/zylog-demo

Configuration

To configure the logger, simply pass an options struct reference to
SetupLogging. For example,

package main

	import (
		logger "github.com/zylisp/zylog/logger"
		log "github.com/sirupsen/logrus"
	)

	func main () {
		log.SetupLogging(&log.ZyLogOptions{
			Colored:      true,
			Level:        "info",
			Output:       "stdout",
			ReportCaller: false,
		})
		// More app code
		log.Info("App started up!")
	}

*/
package logger

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	log "github.com/sirupsen/logrus"
)

// TextFormatter formats logs into text.
type TextFormatter struct {
	// Force disabling colors.
	DisableColors bool
}

// The Options used by the zylog logger to set up logrus.
type ZyLogOptions struct {
	Colored      bool
	Level        string
	Output       string // stdout, stderr, or filesystem
	ReportCaller bool
}

const (
	LogLevelError       string = "Could not set configured log level"
	LogOutputError      string = "Unsupported log output: %s"
	NotImplementedError string = "Not yet implemented: %s"
)

// Logger setup function.
func SetupLogging(opts *ZyLogOptions) {
	level, err := log.ParseLevel(opts.Level)
	if err != nil {
		panic(LogLevelError)
	}
	log.SetLevel(level)
	switch opts.Output {
	case "stdout":
		log.SetOutput(os.Stdout)
	case "stderr":
		log.SetOutput(os.Stderr)
	case "filesystem":
		panic(fmt.Sprintf(NotImplementedError, "filesystem log output"))
	default:
		panic(fmt.Sprintf(LogOutputError, opts.Output))
	}
	disableColors := !opts.Colored
	color.NoColor = disableColors
	log.SetFormatter(&TextFormatter{
		DisableColors: disableColors,
	})
	log.SetReportCaller(opts.ReportCaller)
	log.Info("Logging initialized.")
}

// Provides the custom formatting of the zylog logger.
//
// In particular, logs output in the following form:
//
//	YYYY-mm-DDTHH:MM:SS-TZ:00 LEVEL ▶ logged message ...
//
// If the ReportCaller option is set to true, the log output will have the
// following form:
//
//	YYYY-mm-DDTHH:MM:SS-TZ:00 LEVEL [pkghost/auth/proj/file.Func:LINENUM] ▶ logged message ...
//
// Any structured data passed as logrus fields will be appended to the above
// line forms.
func (f *TextFormatter) Format(entry *log.Entry) ([]byte, error) {
	var b *bytes.Buffer

	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	time := color.GreenString(entry.Time.Format(time.RFC3339))
	level := ColorLevel(strings.ToUpper(entry.Level.String()))

	b.WriteString(fmt.Sprintf("%s %s", time, level))
	if entry.Logger.ReportCaller {
		b.WriteString(fmt.Sprintf(" [%s:%s]",
			color.HiYellowString(entry.Caller.Function),
			color.YellowString(strconv.Itoa(entry.Caller.Line))))
	}
	if entry.Message != "" {
		b.WriteString(color.CyanString(" ▶ "))
		b.WriteString(entry.Message)
	}

	if len(entry.Data) > 0 {
		b.WriteString(" || ")
	}
	for key, value := range entry.Data {
		b.WriteString(fmt.Sprintf("%s={%s}, ", key, value))
	}

	b.WriteByte('\n')
	return b.Bytes(), nil
}

// Determine the color of the log level based upon the string value of the log
// level.
func ColorLevel(level string) string {
	switch level {
	case "TRACE":
		level = color.HiMagentaString(level)
	case "DEBUG":
		level = color.HiCyanString(level)
	case "INFO":
		level = color.HiGreenString(level)
	case "WARNING":
		level = color.HiYellowString(level)
	case "ERROR":
		level = color.RedString(level)
	case "FATAL":
		level = color.HiRedString(level)
	case "PANIC":
		level = color.HiWhiteString(level)
	}
	return level
}
