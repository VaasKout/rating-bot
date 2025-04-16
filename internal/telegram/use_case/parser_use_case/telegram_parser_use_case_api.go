package parser_use_case

import (
	"rating-bot/internal/layout_parser"
)

type TelegramParserUseCaseApi interface {
	GetAppInfo(packageName string) *layout_parser.MarketApp
}
