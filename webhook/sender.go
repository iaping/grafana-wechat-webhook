package webhook

import (
	"errors"
	"fmt"

	"github.com/iaping/go-wechat-robot/robot"
	"github.com/iaping/go-wechat-robot/robot/message"
)

type Sender struct {
	Wechat map[string]*robot.Robot
}

// 发送通知
func (s *Sender) Send(channel, key string, notify *Notification) error {
	switch channel {
	case "wechat":
		return s.sendWechat(key, notify)
	}

	return fmt.Errorf("not found channel [%s]", channel)
}

// 发送企业微信机器人通知
func (s *Sender) sendWechat(key string, notify *Notification) error {
	robot, found := s.Wechat[key]
	if !found {
		return fmt.Errorf("not found wechat robot [%s]", key)
	}

	msg := message.NewMarkdownSimple(notify.BuildMarkdownTemplate())

	resp, err := robot.Send(msg)
	if err != nil {
		return err
	}

	if !resp.IsSuccess() {
		return errors.New(resp.Message)
	}

	return nil
}
