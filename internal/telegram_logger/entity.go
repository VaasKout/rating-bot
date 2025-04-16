package telegram_logger

import (
	"fmt"
	"time"
)

const USER_LOG_FORMAT = "<b>user_id:</b> %d\n" +
	"<b>tag:</b> @%s\n" +
	"<b>message:</b> %s"

const BOT_LOG_FORMAT = "<b>bot_answer:\n</b>%s"

func GetUserMessage(message string, userId int64, userTag string) string {
	if message == "" {
		return ""
	}
	outputMessage := fmt.Sprintf(USER_LOG_FORMAT, userId, userTag, message)
	messageWithDate := outputMessage + fmt.Sprintf("\n<b>date:</b> %s", getTimeDate())
	return messageWithDate
}

func GetBotAnswerMessage(message string) string {
	if message == "" {
		return ""
	}
	outputMessage := fmt.Sprintf(BOT_LOG_FORMAT, message)
	messageWithDate := outputMessage + fmt.Sprintf("\n<b>date:</b> %s", getTimeDate())
	return messageWithDate
}

func getTimeDate() string {
	return time.Now().Format("15:04:05 02-01-06")
}
