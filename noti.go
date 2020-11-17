package noti

import (
	"fmt"
	"strings"

	"github.com/hack-fan/config"
	"go.uber.org/zap"
)

// All sender provider
const (
	ProviderWework = "wework"
)

// Settings will be defined in New func or from Env
type Settings struct {
	NotiProvider string `default:"debug"`
	Wework       WeworkSender
}

// Sender interface
type Sender interface {
	Info(args ...interface{}) error
	Warn(args ...interface{}) error
	Error(args ...interface{}) error
	InfoMD(lines []string) error
	WarnMD(lines []string) error
	ErrorMD(lines []string) error
}

// Noti is noti instance
type Noti struct {
	sender   Sender
	settings *Settings
	log      *zap.SugaredLogger
}

var defaultNoti *Noti

func NewNoti(settings *Settings) *Noti {
	// load more settings from env
	config.MustLoad(settings)
	// new noti
	var n = &Noti{
		settings: settings,
	}
	// provider
	var warning string
	switch settings.NotiProvider {
	case ProviderWework:
		if settings.Wework.Ready() {
			n.sender = settings.Wework
		} else {
			warning = fmt.Sprintf("%s sender config is invalid, check it please", settings.NotiProvider)
		}
	}
	// logger
	if n.sender != nil {
		logger, _ := zap.NewProduction()
		n.log = logger.Sugar()
	} else {
		logger, _ := zap.NewDevelopment()
		n.log = logger.Sugar()
	}

	if warning != "" {
		n.log.Error(warning)
	}

	return n
}

// SetDebug force set debug mode
func (n *Noti) SetDebug() {
	n.sender = nil
	logger, _ := zap.NewDevelopment()
	n.log = logger.Sugar()
}

// SetLogger accept a custom zap sugared logger
func (n *Noti) SetLogger(logger *zap.SugaredLogger) {
	n.log = logger
}

func (n *Noti) Error(args ...interface{}) {
	if n.sender != nil {
		err := n.sender.Error(args...)
		if err != nil {
			n.log.Errorf("send notification to %s failed:%s", n.settings.NotiProvider, err)
		}
	}
	n.log.Error(args...)
}

func (n *Noti) Errorf(format string, a ...interface{}) {
	n.Error(fmt.Sprintf(format, a...))
}

func (n *Noti) ErrorMD(lines []string) {
	if n.sender != nil {
		err := n.sender.ErrorMD(lines)
		if err != nil {
			n.log.Errorf("send notification to %s failed:%s", n.settings.NotiProvider, err)
		}
	}
	n.log.Error(strings.Join(lines, "\n"))
}

func (n *Noti) Warn(args ...interface{}) {
	if n.sender != nil {
		err := n.sender.Warn(args...)
		if err != nil {
			n.log.Errorf("send notification to %s failed:%s", n.settings.NotiProvider, err)
		}
	}
	n.log.Warn(args...)
}

func (n *Noti) Warnf(format string, a ...interface{}) {
	n.Warn(fmt.Sprintf(format, a...))
}

func (n *Noti) WarnMD(lines []string) {
	if n.sender != nil {
		err := n.sender.WarnMD(lines)
		if err != nil {
			n.log.Errorf("send notification to %s failed:%s", n.settings.NotiProvider, err)
		}
	}
	n.log.Warn(strings.Join(lines, "\n"))
}

func (n *Noti) Info(args ...interface{}) {
	if n.sender != nil {
		err := n.sender.Info(args...)
		if err != nil {
			n.log.Errorf("send notification to %s failed:%s", n.settings.NotiProvider, err)
		}
	}
	n.log.Info(args...)
}

func (n *Noti) Infof(format string, a ...interface{}) {
	n.Info(fmt.Sprintf(format, a...))
}

func (n *Noti) InfoMD(lines []string) {
	if n.sender != nil {
		err := n.sender.InfoMD(lines)
		if err != nil {
			n.log.Errorf("send notification to %s failed:%s", n.settings.NotiProvider, err)
		}
	}
	n.log.Info(strings.Join(lines, "\n"))
}

// init default noti, easy for use
func init() {
	var settings = new(Settings)
	defaultNoti = NewNoti(settings)
}

// SetDebug set default noti to debug mode
func SetDebug() {
	defaultNoti.SetDebug()
}

// Error send default error notification
func Error(args ...interface{}) {
	defaultNoti.Error(args...)
}

// Errorf send default error notification with format
func Errorf(format string, a ...interface{}) {
	defaultNoti.Errorf(format, a...)
}

// ErrorMD send default markdown error notification
func ErrorMD(lines []string) {
	defaultNoti.ErrorMD(lines)
}

// Warn send default warn notification
func Warn(args ...interface{}) {
	defaultNoti.Warn(args...)
}

// Warnf send default warn notification with format
func Warnf(format string, a ...interface{}) {
	defaultNoti.Warnf(format, a...)
}

// WarnMD send default markdown warn notification
func WarnMD(lines []string) {
	defaultNoti.WarnMD(lines)
}

// Info send default info notification
func Info(args ...interface{}) {
	defaultNoti.Info(args...)
}

// Infof send default info notification with format
func Infof(format string, a ...interface{}) {
	defaultNoti.Infof(format, a...)
}

// InfoMD send default markdown info notification
func InfoMD(lines []string) {
	defaultNoti.InfoMD(lines)
}
