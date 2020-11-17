package noti

import (
	"fmt"
	"strings"

	"github.com/levigross/grequests"
)

// WeworkSender can send notification to wechat work.
type WeworkSender struct {
	BaseURL  string `default:"https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key="`
	InfoKey  string
	WarnKey  string
	ErrorKey string
}

// RobotMsg is message api model
type RobotMsg struct {
	MsgType  string       `json:"msgtype"`
	Text     *MsgText     `json:"text,omitempty"`
	MarkDown *MsgMarkdown `json:"markdown,omitempty"`
}

// MsgText is text message api model
type MsgText struct {
	Content string `json:"content"`
}

// MsgMarkdown is md message api model
type MsgMarkdown struct {
	Content string `json:"content"`
}

// Ready check if wechat work sender ready
func (s WeworkSender) Ready() bool {
	if s.InfoKey != "" && s.WarnKey != "" && s.ErrorKey != "" {
		return true
	}
	return false
}

// SendRobotMsg send robot message by wechat work web api
func (s WeworkSender) SendRobotMsg(key, tp, content string) error {
	var msg = &RobotMsg{
		MsgType: tp,
	}
	switch tp {
	case "text":
		msg.Text = &MsgText{Content: content}
	case "markdown":
		msg.MarkDown = &MsgMarkdown{Content: content}
	}
	_, err := grequests.Post(s.BaseURL+key, &grequests.RequestOptions{
		JSON: msg,
	})
	if err != nil {
		return fmt.Errorf("wechat work send robot message api error: %w", err)
	}
	return nil
}

// SendRobotMarkdown 向机器人发送 Markdown 通知
func (s WeworkSender) SendRobotMarkdown(key string, lines []string) error {
	// 一次最多4096字 控制一下
	var buffer = make([]string, 0)
	var count, times int
	for _, line := range lines {
		if len(line) > 4000 {
			return fmt.Errorf("markdown message line length must less than 4000")
		}
		if count+len(line) > 4000 {
			err := s.SendRobotMsg(key, "markdown", strings.Join(buffer, "\n"))
			if err != nil {
				return err
			}
			buffer = make([]string, 0)
			count = 0
			times++
			if times >= 5 {
				return fmt.Errorf("markdown message is longger than 20k")
			}
		}
		buffer = append(buffer, line)
		count += len(line)
	}
	// send left in buffer
	return s.SendRobotMsg(key, "markdown", strings.Join(buffer, "\n"))
}

// =============== 三个通知机器人 出错，紧急和普通 ===================

// Error 程序出错通知
func (s WeworkSender) Error(args ...interface{}) error {
	return s.SendRobotMsg(s.ErrorKey, "text", fmt.Sprint(args...))
}

// Warn 紧急通知
func (s WeworkSender) Warn(args ...interface{}) error {
	return s.SendRobotMsg(s.WarnKey, "text", fmt.Sprint(args...))
}

// Info 一般通知
func (s WeworkSender) Info(args ...interface{}) error {
	return s.SendRobotMsg(s.InfoKey, "text", fmt.Sprint(args...))
}

// ErrorMD 出错通知 Markdown
func (s WeworkSender) ErrorMD(lines []string) error {
	return s.SendRobotMarkdown(s.ErrorKey, lines)
}

// WarnMD 紧急通知 Markdown
func (s WeworkSender) WarnMD(lines []string) error {
	return s.SendRobotMarkdown(s.WarnKey, lines)
}

// InfoMD 一般通知 Markdown
func (s WeworkSender) InfoMD(lines []string) error {
	return s.SendRobotMarkdown(s.InfoKey, lines)
}
