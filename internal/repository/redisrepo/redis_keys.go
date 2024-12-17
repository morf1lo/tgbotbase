package redisrepo

import "fmt"

var (
	userKey = "user:%d" // <telegramID>
)

func UserKey(telegramID int64) string {
	return fmt.Sprintf(userKey, telegramID)
}
