package utils

import (
	newraccoon "echo-discord-bot/utils/new_raccoon"

	"github.com/bwmarrin/discordgo"
)

func Utils(s *discordgo.Session) {
	newraccoon.Hello(s)
}
