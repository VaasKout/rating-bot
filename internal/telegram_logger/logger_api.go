package telegram_logger

type LoggerApi interface {
	EnqueueBotLogMessage(message string)
	EnqueueUserBotLogMessage(message string, userId int64, userTag string)
	SendLogs()
}
