package redis

type RedisApi interface {
	SetData(key string, value string) error
	GetData(key string) string
	DeleteData(key string) error
	SAdd(key string, member string) error
	SMembers(key string) []string
	SISMembers(key string, value string) bool
	SRem(key string, member string) error
	RPush(key string, value string) error
	LPush(key string, value string) error
	LPop(key string) string
	GetSize(key string) int64
	LTrim(key string, startIndex int64) error
	LRange(key string) []string
}
