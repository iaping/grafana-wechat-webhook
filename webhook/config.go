package webhook

import "github.com/iaping/go-wechat-robot/robot"

type Config struct {
	Addr   string
	Wechat map[string]string
	Ding   map[string]string
}

func (c Config) WechatRobot() map[string]*robot.Robot {
	robots := map[string]*robot.Robot{}

	for key, hook := range c.Wechat {
		robots[key] = robot.New(hook)
	}

	return robots
}
