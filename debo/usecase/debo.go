package usecase

import (
	"fmt"

	"github.com/bwmarrin/discordgo"

	"github.com/hiroju/discord-botkun/debo/usecase/handlers"
)

type (
	Debo struct {
		Token string
		Name  string
		Dg    *discordgo.Session
	}
)

var (
	TOKEN   = ""
	NAME    = ""
	DEBOKUN Debo
)

func GetDebo() (Debo, error) {
	if (DEBOKUN != Debo{}) {
		return DEBOKUN, nil
	}

	debo, err := initDebo()
	if err != nil {
		return Debo{}, err
	}
	DEBOKUN = debo
	return DEBOKUN, nil
}

func initDebo() (Debo, error) {
	debo, err := createDebo(TOKEN, NAME)
	if err != nil {
		return Debo{}, nil
	}
	debo.AddHandler(handlers.MessageCreate)
	debo.AddHandler(handlers.GetAnimeInfo)
	return debo, nil
}

func createDebo(token string, name string) (Debo, error) {
	dg, err := discordgo.New("Bot " + TOKEN)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return Debo{}, err
	}
	return Debo{Token: token, Name: name, Dg: dg}, nil
}

func (debo *Debo) getDiscordClient() *discordgo.Session {
	return debo.Dg
}

func (debo *Debo) AddHandler(handler interface{}) {
	debo.getDiscordClient().AddHandler(handler)
}

func (debo *Debo) Open() error {
	dg := debo.getDiscordClient()
	err := dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return err
	}
	return nil
}

func (debo *Debo) Close() {
	debo.getDiscordClient().Close()
}
