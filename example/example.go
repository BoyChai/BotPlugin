package main

import (
	"github.com/BoyChai/BotPlugin"
	"github.com/BoyChai/CoralBot"
)

func main() {
	e := BotPlugin.GetEvent()
	h := BotPlugin.GetHandler()
	//设置插件信息
	BotPlugin.SetInfo(CoralBot.PluginInfo{
		Name:      "demo",
		Summary:   "测试插件demo",
		Version:   "v0.0.1",
		Developer: "BoyChai",
		Email:     "1972567225@qq.com",
	})
	BotPlugin.NewTask(CoralBot.Task{
		Condition: []CoralBot.Condition{
			{
				Key:   &e.Message,
				Value: "插件测试",
			},
		},
		RunName: "demo",
	})

	BotPlugin.NewHandles("demo", func(event *CoralBot.Event) error {
		_, err := h.Reply(*e, CoralBot.Msg{Message: "测试成功"})
		if err != nil {
			return err
		}
		return nil
	})
	BotPlugin.BuildPlugin()
}
