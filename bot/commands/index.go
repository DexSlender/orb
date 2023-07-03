package commands

import (
	"github.com/dexslender/orb/bot/util"
)

var (
	commands = []util.Command{
		new(PingCommand),
	}
	Handler util.Handler
)

func init() {
	Handler = util.Handler{}
	Handler.LoadCommands(commands)
}