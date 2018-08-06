package handlers

import (
	"github.com/bwmarrin/discordgo"
)

func IsBotMessage(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	if m.Author.ID == s.State.User.ID {
		return true
	}
	return false
}
