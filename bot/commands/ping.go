package commands

import (
	"github.com/dexslender/orb/bot/util"
	"github.com/disgoorg/disgo/discord"
)

type PingCommand struct {
	discord.SlashCommandCreate
}

func (c PingCommand) Init() {
	c.Name = "Ping"
	c.Description = "Get the latency of this bot"
}

func (c PingCommand) Run(ctx util.CommandPayload) error {
	return nil
}
