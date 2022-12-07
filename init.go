package BotPlugin

import (
	"github.com/BoyChai/CoralBot"
)

type Handle struct {
	name string
	run  func(event *CoralBot.Event) error
}

type config struct {
	Handler []handler `yaml:"handler"`
}

type handler struct {
	Name      string `yaml:"name"`
	Host      string `yaml:"host"`
	Agreement string `yaml:"agreement"`
}

var event CoralBot.Event

var han CoralBot.Handle

var information CoralBot.PluginInfo

type Plugin struct{}

var handles []Handle
