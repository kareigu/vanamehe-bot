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
		&joinInfo,
		&leaveInfo,
	}

	Handlers = map[string]func(
		s *discordgo.Session,
		i *discordgo.InteractionCreate){
		huuttisInfo.Name:    huuttisCmd,
		jattisInfo.Name:     jattisCmd,
		kaannosInfo.Name:    kaannosCmd,
		kiinnostaaInfo.Name: kiinnostaaCmd,
		joinInfo.Name:       joinCmd,
		leaveInfo.Name:      leaveCmd,
	}
)

func InteractionRespondMessage(s *discordgo.Session, i *discordgo.InteractionCreate, msg string) error {
	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionApplicationCommandResponseData{
			Content: msg,
		},
	})
}
