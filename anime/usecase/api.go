package usecase

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/hiroju/discord-botkun/model"
)

var (
	AnimeAPIURL     = "http://api.moemoe.tokyo/anime/v1/master"
	MonthSeasonList = map[string]string{
		"1":  "1",
		"2":  "1",
		"3":  "1",
		"4":  "2",
		"5":  "2",
		"6":  "2",
		"7":  "3",
		"8":  "3",
		"9":  "3",
		"10": "4",
		"11": "4",
		"12": "4",
	}
)

func GetAnimeInfoYearSeason(year string, season string) (model.AnimeInfoYearSeasons, error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", getURL(year, season), nil)
	if err != nil {
		// handle error
		fmt.Fprintln(os.Stdout, err)
		return model.AnimeInfoYearSeasons{}, err
	}
	resp, err := client.Do(req)
	if err != nil {
		// handle error
		fmt.Fprintln(os.Stdout, err)
		return model.AnimeInfoYearSeasons{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Fprintln(os.Stdout, err)
		return model.AnimeInfoYearSeasons{}, err
	}

	// jsonBytes := ([]byte)(jsonStr)
	jsonBytes := body
	data := model.AnimeInfoYearSeasons{}

	if err := json.Unmarshal(jsonBytes, &data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return model.AnimeInfoYearSeasons{}, nil
	}

	fmt.Fprintln(os.Stdout, data)
	return data, nil
}

func ConvertToDiscordFormat(infos model.AnimeInfoYearSeasons) string {
	msg := ""
	for i := 0; i < len(infos); i++ {
		info := infos[i]
		msg += info.Title + "\t" + info.PublicURL + "\n"
	}
	return msg
}

func ConvertMonthToSeason(month string) string {
	if _, ok := MonthSeasonList[month]; !ok {
		return ""
	}
	return MonthSeasonList[month]
}

func getURL(year string, season string) string {
	return AnimeAPIURL + "/" + year + "/" + season
}
