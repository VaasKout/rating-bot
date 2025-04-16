package app

import (
	"rating-bot/configs"
	"rating-bot/internal/data/telegram_redis"
	"rating-bot/internal/layout_parser"
	"rating-bot/internal/telegram/controller"
	"rating-bot/internal/telegram_logger"
	"rating-bot/pkg/core/redis"
	"rating-bot/pkg/logger"
)

type App struct {
	Config *configs.Config
	Logger *logger.Logger
}

func New() *App {
	cfg := configs.New()
	zapLogger := logger.New()
	return &App{
		Config: cfg,
		Logger: zapLogger,
	}
}

func (app *App) Run() {
	ratingMsgChan := make(chan telegram_redis.Message, 10)

	coreRedis := redis.New(app.Config)
	layoutParser := layout_parser.New(app.Logger, app.Config.RootProps.ConfigPath)

	tgLogger := telegram_logger.New(app.Config, app.Logger, coreRedis)
	tgBot := controller.New(app.Logger, app.Config, coreRedis, layoutParser, ratingMsgChan, tgLogger)

	go tgBot.ListenUpdates()
	go tgLogger.SendLogs()
	tgBot.SendEnqueuedMessages()
}
