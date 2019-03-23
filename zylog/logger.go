package zylog

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

// TextFormatter formats logs into text
type TextFormatter struct {
	// Force disabling colors.
	DisableColors bool
}

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
