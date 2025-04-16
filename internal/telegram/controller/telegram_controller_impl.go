package controller

import (
	"rating-bot/configs"
	"rating-bot/internal/data/telegram_redis"
	"rating-bot/internal/layout_parser"
	"rating-bot/internal/telegram/navigation"
	"rating-bot/internal/telegram/use_case"
	"rating-bot/internal/telegram_logger"
	"rating-bot/pkg/core/redis"
	"rating-bot/pkg/logger"
	"time"
)

type TelegramControllerImpl struct {
	telegramUseCase use_case.TelegramUseCase
	navHandler      navigation.HandlerApi
	layoutParser    layout_parser.LayoutParserApi
	tgLogger        telegram_logger.LoggerApi
}

func New(
	log *logger.Logger,
	cfg *configs.Config,
	coreRds redis.RedisApi,
	layoutParser layout_parser.LayoutParserApi,
	ratingMsgChan chan telegram_redis.Message,
	tgLogger telegram_logger.LoggerApi,
) TelegramControllerApi {
	var telegramUseCase = use_case.New(log, cfg, coreRds, layoutParser, ratingMsgChan)
	var navHandler = navigation.New(telegramUseCase)

	return &TelegramControllerImpl{
		telegramUseCase: telegramUseCase,
		navHandler:      navHandler,
		layoutParser:    layoutParser,
		tgLogger:        tgLogger,
	}
}

func (controller *TelegramControllerImpl) SendEnqueuedMessages() {
	for {
		time.Sleep(time.Millisecond * 100)
		var message = controller.telegramUseCase.RedisUseCase().PopMessage()
		if message == nil || message.Text == "" {
			continue
		}
		controller.tgLogger.EnqueueBotLogMessage(message.Text)
		controller.telegramUseCase.RedisUseCase().SendTelegramMessage(message)
	}
}

func (controller *TelegramControllerImpl) ListenUpdates() {
	for {
		messages := controller.telegramUseCase.RedisUseCase().GetNewMessages()
		for _, item := range *messages {
			controller.tgLogger.EnqueueUserBotLogMessage(item.Text, item.ChatId, item.UserName)
			controller.navHandler.HandleMessage(&item)
		}
		time.Sleep(time.Millisecond * 500)
	}
}
