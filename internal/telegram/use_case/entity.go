package use_case

import (
	"fmt"
	"rating-bot/internal/data/telegram_redis"
	"rating-bot/internal/layout_parser"
	"rating-bot/internal/telegram/domain"
	"sort"
)

var resultAppInfoMessage = func(app *layout_parser.MarketApp) *telegram_redis.Message {
	var msg = ""
	if app == nil || app.Title == "" || app.Developer == "" {
		msg = domain.USER_GET_RATING_APP_NOT_FOUND
	} else {
		rating := ""
		if len(app.Rating) == 0 {
			rating = domain.USER_GET_RATING_NO_RATING
		} else {
			ratingArr := make([]string, 0)
			for _, item := range app.Rating {
				ratingArr = append(ratingArr, fmt.Sprintf("%s: <b>%s</b>\n", item.CountryName, item.Rating))
			}
			sort.Strings(ratingArr)
			for _, item := range ratingArr {
				rating += item
			}
		}

		msg = fmt.Sprintf(
			domain.USER_GET_RATING_CONFIRM_APP_MESSAGE,
			app.Title,
			app.Developer,
			app.Package,
			app.FoundUrl,
			rating,
		)
	}
	return telegram_redis.InitOutputMessage(msg, &[][]string{})
}
