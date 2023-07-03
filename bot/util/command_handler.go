package util

import (
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/disgo/handler"
)

var _ = handler.Mux{}

type Command interface {
	discord.ApplicationCommandCreate
	Init()
	Run(CommandPayload) error
}

var _ bot.EventListener = (*Handler)(nil)

type Handler struct {
	commands []Command
}

func (h *Handler) LoadCommands(c []Command) {
}

func (h *Handler) Get(name string) (ok bool, command Command) {
	for _, c := range h.commands {
		if c.CommandName() == name {
			ok, command = true, c
			break
		}
	}
	return
}

func (h *Handler) ApplicationCommand() (commands []discord.ApplicationCommandCreate) {
	for _, c := range h.commands {
		commands = append(commands, discord.ApplicationCommandCreate(c))
	}
	return
}

func (h *Handler) OnEvent(event bot.Event) {
	e, ok := event.(events.InteractionCreate)

	// e, ok := event.(events.InteractionCreate)
	if !ok {
		return
	}

	switch i := e.Interaction.(type) {
	case discord.ApplicationCommandInteraction:
		if ok, c := h.Get(i.Data.CommandName()); ok {
			c.Run(CommandPayload{
				ApplicationCommandInteractionCreate: events.ApplicationCommandInteractionCreate{
					e.GenericEvent,
					e.Interaction.(discord.ApplicationCommandInteraction),
					e.Respond,
				},
			})
		}
	case discord.AutocompleteInteraction:
	case discord.ModalSubmitInteraction:
	case discord.ComponentInteraction:
	}
}

// Represents the events.ApplicationCommandInteractionCreate event
type CommandPayload struct {
	events.ApplicationCommandInteractionCreate
}

type ModalPayload struct {
}
