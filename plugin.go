package BotPlugin

import (
	"github.com/BoyChai/CoralBot"
)

func (p *Plugin) PluginInfo(n string, info *CoralBot.PluginInfo) error {
	*info = information
	h := GetHandler()
	c, err := readConfig()
	if err != nil {
		return err
	}
	h.Host = c.Host
	h.Agreement = c.Agreement
	return nil
}

func (p *Plugin) GetPlugin(e *CoralBot.Event, Task *[]CoralBot.Task) error {
	event = *e
	*Task = CoralBot.Tasks
	return nil
}

func (p *Plugin) Handles(e *CoralBot.Event, n *string) error {
	for _, value := range handles {
		if e.Other.RunName == value.name {
			err := value.run(e)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
