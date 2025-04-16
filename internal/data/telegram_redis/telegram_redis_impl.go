package telegram_redis

import (
	"fmt"
	"rating-bot/pkg/core/redis"
	"strconv"
)

const USERS_KEY = "users"
const ADMINS_KEY = "admins"
const SUPERVISORS_KEY = "supervisors"
const SEND_MESSAGE_QUEUE_KEY = "send_message_queue"
const TELEGRAM_OFFSET_KEY = "telegram_offset"

type TelegramRedisImpl struct {
	redis redis.RedisApi
}

func New(coreRedis redis.RedisApi) TelegramRedisApi {
	return &TelegramRedisImpl{
		redis: coreRedis,
	}
}

func (adapter TelegramRedisImpl) IsUser(userName string) bool {
	return adapter.redis.SISMembers(USERS_KEY, userName)
}

func (adapter TelegramRedisImpl) IsAdmin(userName string) bool {
	return adapter.redis.SISMembers(ADMINS_KEY, userName)
}

func (adapter TelegramRedisImpl) IsSupervisor(userName string) bool {
	return adapter.redis.SISMembers(SUPERVISORS_KEY, userName)
}

func (adapter TelegramRedisImpl) SaveUserData(userData *UserData) bool {
	var userJson = MapUserDataToJson(userData)
	err := adapter.redis.SetData(fmt.Sprint(userData.ChatId), userJson)
	return err == nil
}

func (adapter TelegramRedisImpl) ClearUserData(userId int64) bool {
	err := adapter.redis.DeleteData(fmt.Sprint(userId))
	return err == nil
}

func (adapter TelegramRedisImpl) GetUserData(userId int64) *UserData {
	result := adapter.redis.GetData(fmt.Sprint(userId))
	userData := MapJsonToUserData(result)
	return userData
}

func (adapter TelegramRedisImpl) RPushMessage(message *Message) bool {
	var msgJson = MapMessageToJson(message)
	err := adapter.redis.RPush(SEND_MESSAGE_QUEUE_KEY, msgJson)
	return err == nil
}

func (adapter TelegramRedisImpl) LPushMessage(message *Message) bool {
	var msgJson = MapMessageToJson(message)
	err := adapter.redis.LPush(SEND_MESSAGE_QUEUE_KEY, msgJson)
	return err == nil
}

func (adapter TelegramRedisImpl) PopMessage() *Message {
	result := adapter.redis.LPop(SEND_MESSAGE_QUEUE_KEY)
	if len(result) == 0 {
		return nil
	}
	return MapJsonToMessage(result)
}

func (adapter TelegramRedisImpl) SaveOffset(value int64) bool {
	err := adapter.redis.SetData(TELEGRAM_OFFSET_KEY, fmt.Sprint(value))
	return err == nil
}

func (adapter TelegramRedisImpl) GetOffset() int64 {
	result := adapter.redis.GetData(TELEGRAM_OFFSET_KEY)
	i, err := strconv.ParseInt(result, 10, 64)
	if err != nil {
		return 0
	}
	return i
}
