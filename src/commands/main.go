package commands

import (
	"github.com/bwmarrin/discordgo"
)

var (
	List = []*discordgo.ApplicationCommand{
		&huuttisInfo,
		&jattisInfo,
		&kaannosInfo,
		&kiinnostaaInfo,
	}

	Handlers = map[string]func(
		s *discordgo.Session,
		i *discordgo.InteractionCreate){
		huuttisInfo.Name:    huuttisCmd,
		jattisInfo.Name:     jattisCmd,
		kaannosInfo.Name:    kaannosCmd,
		kiinnostaaInfo.Name: kiinnostaaCmd,
	}
)
