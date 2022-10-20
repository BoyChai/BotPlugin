package BotPlugin

import "github.com/BoyChai/CoralBot"

var Info CoralBot.PluginInfo

type MyPlugin struct{}

func (p *MyPlugin) PluginInfo(n interface{}, info *CoralBot.PluginInfo) error {
	*info = Info
	return nil
}
