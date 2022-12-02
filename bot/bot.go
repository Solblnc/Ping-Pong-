package bot

import (
	"DISCORD-PING-BOT/config"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

var BotID string

//var Botting *discordgo.Session

func Start() {
	Botting, err := discordgo.New("bot " + config.Token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	u, err := Botting.User("@me")

	if err != nil {
		fmt.Println(err.Error())

	}

	u.ID = BotID

	Botting.AddHandler(messageHandler)

	err = Botting.Open()
	if err != nil {
		fmt.Println(err.Error())
		return

	}
	fmt.Println("Bot is running")

}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID != BotID {
		return
	}
	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
	}
}
