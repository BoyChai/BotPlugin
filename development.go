package BotPlugin

import (
	"errors"
	"github.com/BoyChai/CoralBot"
	"github.com/dullgiulio/pingo"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

func readConfig() (handler, error) {
	var c config
	var all handler
	var my handler
	var n handler
	data, err := ioutil.ReadFile("./plugin/config.yaml")
	if err != nil {
		return n, err
	}
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		return handler{}, err
	}
	for _, value := range c.Handler {
		if value.Name == information.Name {
			my = value
			continue
		}
		if value.Name == "all" {
			all = value
			continue
		}
	}

	if my != n {
		return my, nil
	}
	if all != n {
		return all, nil
	}
	return handler{}, errors.New("未读取到配置文件")
}

func GetEvent() *CoralBot.Event {
	return &event
}

func GetHandler() *CoralBot.Handle {
	return &han
}

//func GetHandlers() []Handle {
//	return handles
//}

func SetInfo(info CoralBot.PluginInfo) {
	information = info
}

func NewTask(task CoralBot.Task) {
	CoralBot.NewPluginTask(task)
	CoralBot.Tasks = append(CoralBot.Tasks, task)
}

func NewHandles(name string, run func(event *CoralBot.Event) error) {
	handles = append(handles, Handle{
		name: name,
		run:  run,
	})
}

func BuildPlugin() {
	plugin := &Plugin{}
	pingo.Register(plugin)
	pingo.Run()
}
