package parser_use_case

import (
	"rating-bot/internal/layout_parser"
	"strings"
)

type TelegramParserUseCaseImpl struct {
	layoutParser layout_parser.LayoutParserApi
}

func NewLayoutParserUseCase(parser layout_parser.LayoutParserApi) TelegramParserUseCaseApi {
	return &TelegramParserUseCaseImpl{
		layoutParser: parser,
	}
}

func (useCase *TelegramParserUseCaseImpl) GetAppInfo(packageName string) *layout_parser.MarketApp {
	return useCase.layoutParser.GetAppInfo(strings.ToLower(packageName))
}
