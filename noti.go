package noti

import (
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"
)

var debug bool

// 企业微信机器人的 key
var (
	infoKey  string
	warnKey  string
	errorKey string
)

var log *zap.SugaredLogger

// 注意这要在程序运行前存在环境变量才有效
func init() {
	de := os.Getenv("APP_DEBUG")
	if de == "true" || de == "TRUE" || de == "True" || de == "1" {
		SetDebug()
	}

	logger, _ := zap.NewProduction()
	log = logger.Sugar()

	infoKey = os.Getenv("IM_INFO_KEY")
	warnKey = os.Getenv("IM_WARN_KEY")
	errorKey = os.Getenv("IM_ERROR_KEY")
	// 只要配置不齐，就设置成调试模式
	if infoKey == "" || warnKey == "" || errorKey == "" {
		SetDebug()
	}
}

// SetDebug 设置为调试模式，通知打印到日志
func SetDebug() {
	debug = true
}

// Error 企业微信错误通知,调试模式下只打日志
func Error(args ...interface{}) {
	if debug {
		log.Error(args...)
	} else {
		wError(args...)
	}
}

// Errorf 企业微信错误通知,调试模式下只打日志
func Errorf(format string, a ...interface{}) {
	if debug {
		log.Errorf(format, a...)
	} else {
		wError(fmt.Sprintf(format, a...))
	}
}

// ErrorMD Markdown 格式企业微信出错通知,调试模式下只打日志
func ErrorMD(lines []string) {
	if debug {
		log.Info(strings.Join(lines, "\n"))
	} else {
		wErrorMD(lines)
	}
}

// Warn 企业微信重要通知,调试模式下只打日志
func Warn(args ...interface{}) {
	if debug {
		log.Warn(args...)
	} else {
		wWarn(args...)
	}
}

// Warnf 企业微信重要通知,调试模式下只打日志
func Warnf(format string, a ...interface{}) {
	if debug {
		log.Warnf(format, a...)
	} else {
		wWarn(fmt.Sprintf(format, a...))
	}
}

// WarnMD Markdown 格式企业微信重要通知,调试模式下只打日志
func WarnMD(lines []string) {
	if debug {
		log.Warn(strings.Join(lines, "\n"))
	} else {
		wWarnMD(lines)
	}
}

// Info 企业微信通知,调试模式下只打日志
func Info(args ...interface{}) {
	if debug {
		log.Info(args...)
	} else {
		wInfo(args...)
	}
}

// Infof 企业微信通知,调试模式下只打日志
func Infof(format string, a ...interface{}) {
	if debug {
		log.Infof(format, a...)
	} else {
		wInfo(fmt.Sprintf(format, a...))
	}
}

// InfoMD Markdown 格式企业微信一般通知,调试模式下只打日志
func InfoMD(lines []string) {
	if debug {
		log.Info(strings.Join(lines, "\n"))
	} else {
		wInfoMD(lines)
	}
}
