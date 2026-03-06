package utils

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"processBlueprint/internal/config"
)

// ---------- level ----------

type Level int

const (
	LevelError Level = iota + 1
	LevelSuc
	LevelWarn
	LevelInfo
	LevelDebug
)

// ---------- color ----------

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorGray   = "\033[90m"
)

// ---------- prefix ----------

const (
	prefixError = "[error] "
	prefixSuc   = "[suc] "
	prefixWarn  = "[warn] "
	prefixInfo  = "[info] "
	prefixDebug = "[debug] "
)

// ---------- logger ----------

type Logger struct {
	l        *log.Logger
	level    Level
	useColor bool

	task string

	mu sync.Mutex
}

// parseLogLevel converts a string log level to the Level type.
func parseLogLevel(levelStr string) Level {
	switch strings.ToLower(levelStr) {
	case "error":
		return LevelError
	case "suc":
		return LevelSuc
	case "warn":
		return LevelWarn
	case "info":
		return LevelInfo
	case "debug":
		return LevelDebug
	default:
		return LevelInfo // Default to LevelInfo if the input is invalid
	}
}

// ---------- constructor ----------

func NewLogger() *Logger {
	cfg := config.Get()
	return &Logger{
		l:        log.New(os.Stdout, "", log.LstdFlags),
		level:    parseLogLevel(cfg.LogLevel),
		useColor: isTTY(),
	}
}

// 派生一个带 task 的 logger
func (lg *Logger) WithTask(task string) *Logger {
	return &Logger{
		l:        lg.l,
		level:    lg.level,
		useColor: lg.useColor,
		task:     task,
	}
}

// ---------- public api ----------

func (lg *Logger) Fatalf(format string, v ...any) {
	lg.mu.Lock()
	defer lg.mu.Unlock()

	prefixTask := ""
	if lg.task != "" {
		prefixTask = fmt.Sprintf("[%s] ", lg.task)
	}

	msg := fmt.Sprintf(format, v...)

	// 构建完整的日志行
	var logLine string
	if lg.useColor {
		// 先设置颜色，然后打印日志，最后重置颜色
		logLine = fmt.Sprintf("%s%s", colorRed+"[error & exit] "+prefixTask+colorReset, msg)
	} else {
		logLine = fmt.Sprintf("%s%s", "[error & exit] "+prefixTask, msg)
	}

	log.Fatalf(logLine+"%s", "\n")
}

// Error logs a message at error level with format support
// Usage: logger.Error("Failed to process %s: %v", filename, err)
func (lg *Logger) Error(format string, v ...any) {
	lg.print(LevelError, colorRed, prefixError, format, v...)
}

// Suc logs a message at success level with format support
// Usage: logger.Suc("Successfully processed %d items", count)
func (lg *Logger) Suc(format string, v ...any) {
	lg.print(LevelSuc, colorGreen, prefixSuc, format, v...)
}

// Warn logs a message at warning level with format support
// Usage: logger.Warn("Disk space low: %d%% remaining", percent)
func (lg *Logger) Warn(format string, v ...any) {
	lg.print(LevelWarn, colorYellow, prefixWarn, format, v...)
}

// Info logs a message at info level with format support
// Usage: logger.Info("Processing file: %s", filename)
func (lg *Logger) Info(format string, v ...any) {
	lg.print(LevelInfo, colorBlue, prefixInfo, format, v...)
}

// Debug logs a message at debug level with format support
// Usage: logger.Debug("Variable state: %+v", state)
func (lg *Logger) Debug(format string, v ...any) {
	lg.print(LevelDebug, colorGray, prefixDebug, format, v...)
}

// ---------- internal ----------

// print handles the actual logging with proper formatting
func (lg *Logger) print(level Level, color, prefix, format string, v ...any) {
	if level > lg.level {
		// 日志级别未到达要求，不打印
		return
	}
	// Format the message using the format string and arguments
	msg := fmt.Sprintf(format, v...)

	lg.mu.Lock()
	defer lg.mu.Unlock()

	prefixTask := ""
	if lg.task != "" {
		prefixTask = fmt.Sprintf("[%s] ", lg.task)
	}

	// 构建完整的日志行
	var logLine string
	if lg.useColor {
		// 先设置颜色，然后打印日志，最后重置颜色
		logLine = fmt.Sprintf("%s%s", color+prefix+prefixTask+colorReset, msg)
	} else {
		logLine = fmt.Sprintf("%s%s", prefix+prefixTask, msg)
	}

	// 直接打印完整的行，不使用 SetPrefix
	lg.l.Println(logLine)
}

// ---------- legacy support (optional) ----------
// If you want to maintain backward compatibility with the old style,
// you can add these methods:

// Errorln logs a message at error level (legacy support)
func (lg *Logger) Errorln(v ...any) {
	lg.printLegacy(LevelError, colorRed, prefixError, v...)
}

// Sucln logs a message at success level (legacy support)
func (lg *Logger) Sucln(v ...any) {
	lg.printLegacy(LevelSuc, colorGreen, prefixSuc, v...)
}

// Warnln logs a message at warning level (legacy support)
func (lg *Logger) Warnln(v ...any) {
	lg.printLegacy(LevelWarn, colorYellow, prefixWarn, v...)
}

// Infoln logs a message at info level (legacy support)
func (lg *Logger) Infoln(v ...any) {
	lg.printLegacy(LevelInfo, colorBlue, prefixInfo, v...)
}

// Debugln logs a message at debug level (legacy support)
func (lg *Logger) Debugln(v ...any) {
	lg.printLegacy(LevelDebug, colorGray, prefixDebug, v...)
}

// printLegacy handles legacy logging without format string
func (lg *Logger) printLegacy(level Level, color, prefix string, v ...any) {
	if level > lg.level {
		// 日志级别未到达要求，不打印
		return
	}
	msg, fields := splitMessageAndFields(v...)

	lg.mu.Lock()
	defer lg.mu.Unlock()

	prefixTask := ""
	if lg.task != "" {
		prefixTask = fmt.Sprintf("[%s] ", lg.task)
	}

	// 构建完整的日志行
	var logLine string
	if lg.useColor {
		// 先设置颜色，然后打印日志，最后重置颜色
		logLine = fmt.Sprintf("%s%s%s", color+prefix+prefixTask+colorReset, msg, fields)
	} else {
		logLine = fmt.Sprintf("%s%s%s", prefix+prefixTask, msg, fields)
	}

	// 直接打印完整的行，不使用 SetPrefix
	lg.l.Println(logLine)
}

// splitMessageAndFields remains for legacy support
func splitMessageAndFields(v ...any) (string, string) {
	if len(v) == 0 {
		return "", ""
	}

	msg := fmt.Sprint(v[0])

	if len(v) < 3 {
		return msg, ""
	}

	var b strings.Builder
	for i := 1; i+1 < len(v); i += 2 {
		key := fmt.Sprint(v[i])
		val := fmt.Sprint(v[i+1])
		b.WriteString(key)
		b.WriteString("=")
		b.WriteString(val)
		b.WriteString(" ")
	}

	return msg, strings.TrimSpace(b.String())
}

// ---------- tty ----------

func isTTY() bool {
	fi, err := os.Stdout.Stat()
	if err != nil {
		return false
	}
	return (fi.Mode() & os.ModeCharDevice) != 0
}
