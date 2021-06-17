package commands

import (
	"github.com/bwmarrin/discordgo"
)

var (
	List = []*discordgo.ApplicationCommand{
		&huuttisInfo,
	}

	Handlers = map[string]func(
		s *discordgo.Session,
		i *discordgo.InteractionCreate){
		"huuttis": huuttisCmd,
	}
)
