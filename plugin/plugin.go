package plugin

import (
	"log"
	"strings"
)

type Plugin interface {
	Exe(cmd string, args []string) ([]string, error)
}

var (
	plugins = make(map[string]Plugin)
)

func Register(cmd string, p Plugin) {
	cmd = strings.ToLower(cmd)
	if _, ok := Find(cmd); ok {
		log.Printf("plugin.Register, dunplicate of %s\n", cmd)
		return
	}

	plugins[cmd] = p
}

func Find(cmd string) (Plugin, bool) {
	cmd = strings.ToLower(cmd)
	p, ok := plugins[cmd]
	return p, ok
}
