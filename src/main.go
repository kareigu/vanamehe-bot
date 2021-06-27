package main

import (
	"commands"
	"log"
	"os"
	"os/signal"
	"syscall"
	"utils"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	client  *discordgo.Session
	GuildID string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Couldn't load env file")
	}

	bot_token := os.Getenv("TOKEN")
	GuildID = os.Getenv("GUILD_ID")

	if bot_token == "" {
		log.Fatal("Bot token missing from .env")
	}

	client, err = discordgo.New("Bot " + bot_token)
	if err != nil {
		log.Fatal("Couldn't create bot client")
	}
}

func init() {
	client.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commands.Handlers[i.Data.Name]; ok {
			h(s, i)

			var username, id string

			if i.Member != nil {
				username = i.Member.User.Username
				id = i.Member.User.ID
			} else if i.User != nil {
				username = i.User.Username
				id = i.User.ID
			} else {
				log.Printf("Error: no valid user info found")
			}

			log.Printf("%v#%v ran command %v",
				username,
				id,
				i.Data.Name)
		}
	})

	client.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		if m.Author.ID == utils.PRIIDIK_ID && m.Content == "(mis see on)" {
			utils.PlayVoiceLine(s)
		}
	})
}

func main() {
	client.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsGuilds | discordgo.IntentsGuildVoiceStates
	client.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Println("Running...")
	})

	err := client.Open()
	if err != nil {
		log.Fatal("Couldn't start bot client")
	}

	for _, v := range commands.List {
		_, err = client.ApplicationCommandCreate(client.State.User.ID, GuildID, v)
		if err != nil {
			log.Panicf("Couldn't create command %v : %v", v.Name, err)
		}
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	client.Close()
}
