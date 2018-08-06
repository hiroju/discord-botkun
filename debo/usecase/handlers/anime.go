package handlers

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	anime_usecase "github.com/hiroju/discord-botkun/anime/usecase"
	"github.com/hiroju/discord-botkun/command/usecase"
)

var (
	COMMAND_ANIME = "anime"
)

func GetAnimeInfo(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if IsBotMessage(s, m) {
		return
	}

	if !isAnimeCommand(m.Content) {
		return
	}

	args, err := usecase.ParseCommand(COMMAND_ANIME, m.Content)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		return
	}

	if args == nil {
		fmt.Fprintln(os.Stdout, "hoge")
		return
	}

	if len(args) == 0 {
		fmt.Fprintln(os.Stdout, "empty")
	}

	fmt.Fprintln(os.Stdout, args)

	now := time.Now()
	year := ""
	if years, ok := args["y"]; !ok {
		year = now.Format("2006")
	} else {
		year = years[0]
	}

	season := ""
	if seasons, ok := args["s"]; !ok {
		month := now.Format("1")
		season = anime_usecase.ConvertMonthToSeason(month)
	} else {
		season = seasons[0]
	}

	data, err := anime_usecase.GetAnimeInfoYearSeason(year, season)

	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		s.ChannelMessageSend(m.ChannelID, "すまん、エラーでデータとれんかったわ")
	}

	s.ChannelMessageSend(m.ChannelID, anime_usecase.ConvertToDiscordFormat(data))
	return
}

func isAnimeCommand(msg string) bool {
	return strings.HasPrefix(msg, "/"+COMMAND_ANIME)
}
