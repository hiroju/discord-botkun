package model

var (
	AnimeSeasonJA map[string]string = map[string]string{
		"冬": "1",
		"春": "2",
		"夏": "3",
		"秋": "4",
	}
)

type (
	AnimeInfoYearSeasons []AnimeInfoYearSeason
	AnimeInfoYearSeason  struct {
		TitleShort2    string `json:"title_short2"`
		TwitterAccount string `json:"twitter_account"`
		PublicURL      string `json:"public_url"`
		TitleShort1    string `json:"title_short1"`
		Sex            int    `json:"sex"`
		TwitterHashTag string `json:"twitter_hash_tag"`
		ID             int    `json:"id"`
		Sequel         int    `json:"sequel"`
		CreatedAt      string `json:"created_at"`
		CityName       string `json:"city_name"`
		CoursID        int    `json:"cours_id"`
		Title          string `json:"title"`
		CityCode       int    `json:"city_code"`
		TitleShort3    string `json:"title_short3"`
		UpdatedAt      string `json:"updated_at"`
	}
)
