package configs

import (
	"os"
	"strconv"
)

type Config struct {
	RedisProps *RedisProps
	BotProps   *BotProps
	RootProps  *RootProps
}

type RedisProps struct {
	RedisPassword string
	RedisAddress  string
}

type BotProps struct {
	Token        string
	BotLogChatId int64
}

type RootProps struct {
	ConfigPath string
}

func New() *Config {
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisAddress := os.Getenv("REDIS_ADDRESS")

	botToken := os.Getenv("BOT_KEY")
	if botToken == "" {
		panic("Bot key is not found")
	}
	botLogChatID := os.Getenv("BOT_LOG_CHAT_ID")
	if botLogChatID == "" {
		panic("BOT_LOG_CHAT_ID is not found")
	}
	botLogChatIDInt, err := strconv.ParseInt(botLogChatID, 10, 64)
	if err != nil {
		panic("BOT_LOG_CHAT_ID cannot be parsed")
	}

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		panic("CONFIG_PATH not found")
	}

	return &Config{
		RedisProps: &RedisProps{
			RedisPassword: redisPassword,
			RedisAddress:  redisAddress,
		},
		BotProps: &BotProps{
			Token:        botToken,
			BotLogChatId: botLogChatIDInt,
		},
		RootProps: &RootProps{
			ConfigPath: configPath,
		},
	}
}
