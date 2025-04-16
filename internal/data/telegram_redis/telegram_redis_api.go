package telegram_redis

type TelegramRedisApi interface {
	IsUser(userName string) bool
	IsAdmin(userName string) bool
	IsSupervisor(userName string) bool

	SaveUserData(user *UserData) bool
	GetUserData(userId int64) *UserData
	ClearUserData(userId int64) bool

	RPushMessage(message *Message) bool
	LPushMessage(message *Message) bool
	PopMessage() *Message

	SaveOffset(value int64) bool
	GetOffset() int64
}
