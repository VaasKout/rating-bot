package layout_parser

import (
	"encoding/json"
	"fmt"
	"strings"
)

const (
	GOOGLE_PLAY_URL = "https://play.google.com/store/apps/details?id=%s&gl=%s&hl=%s"
	APP_STORE_URL   = "https://apps.apple.com/%s/app/%s"
)

const (
	ITEM_PROP_KEY     = "itemprop"
	NAME_VALUE        = "name"
	STAR_RATING_VALUE = "starRating"
	HREF_KEY          = "href"
	STORE_DEV         = `/store/apps/dev`
	CLASS_KEY         = "class"
	RATING_VALUE      = "we-customer-ratings__averages__display"
	DEVELOPER_VALUE   = "app-privacy__developer-name"
	TITLE_VALUE       = "product-header__title app-header__title"
)

type CountryDB struct {
	CountryCode string `json:"country_code"`
	CountryName string `json:"country_name"`
	LangCode    string `json:"lang_code"`
}

type MarketApp struct {
	Title     string   `json:"title"`
	Developer string   `json:"developer"`
	Package   string   `json:"package"`
	FoundUrl  string   `json:"foundUrl"`
	Rating    []Rating `json:"rating"`
}

type Rating struct {
	CountryName string `json:"country_name"`
	Url         string `json:"url"`
	Rating      string `json:"rating"`
}

func (app *MarketApp) ToJson() string {
	result, err := json.Marshal(app)
	if err != nil {
		fmt.Println("GoogleApplicationToJson")
		fmt.Println(err)
		return ""
	}
	return string(result)
}

func getUrlByPackage(packageName string, countryDb *CountryDB) string {
	if countryDb == nil {
		return ""
	}
	if strings.Contains(packageName, ".") {
		return fmt.Sprintf(GOOGLE_PLAY_URL, packageName, countryDb.CountryCode, countryDb.LangCode)
	}
	if !strings.Contains(packageName, ".") && strings.Contains(packageName, "id") {
		return fmt.Sprintf(APP_STORE_URL, countryDb.CountryCode, packageName)
	}
	return ""
}

func getTitle(htmlResponse string) string {
	title := FindElementInHtml(htmlResponse, ITEM_PROP_KEY, NAME_VALUE)
	if title != "" {
		return strings.TrimSpace(title)
	}
	title = FindElementInHtml(htmlResponse, CLASS_KEY, TITLE_VALUE)
	if title != "" {
		return strings.TrimSpace(title)
	}
	return ""
}

func getDeveloper(htmlResponse string) string {
	developer := FindElementInHtml(htmlResponse, HREF_KEY, STORE_DEV)
	if developer != "" {
		return strings.TrimSpace(developer)
	}
	developer = FindElementInHtml(htmlResponse, CLASS_KEY, DEVELOPER_VALUE)
	if developer != "" {
		return strings.TrimSpace(developer)
	}
	return ""
}

func getRating(htmlResponse string) string {
	rating := FindElementInHtml(htmlResponse, ITEM_PROP_KEY, STAR_RATING_VALUE)
	if rating == "" {
		rating = FindElementInHtml(htmlResponse, CLASS_KEY, RATING_VALUE)
		if rating == "" {
			return ""
		}
	}
	parsedRating := make([]rune, 0)

RatingLoop:
	for _, item := range rating {
		for _, runeMap := range RuneMaps {
			if result, ok := runeMap[item]; ok {
				parsedRating = append(parsedRating, result)
				continue RatingLoop
			}
		}
		parsedRating = append(parsedRating, item)
	}

	return strings.TrimSpace(string(parsedRating))
}
