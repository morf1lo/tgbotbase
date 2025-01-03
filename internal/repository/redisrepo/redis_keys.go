package redisrepo

import "fmt"

var (
	adminKey = "admin:%d" // <telegramID>
	userKey = "user:%d" // <telegramID>
)

func AdminKey(telegramID int64) string {
	return fmt.Sprintf(adminKey, telegramID)
}

func UserKey(telegramID int64) string {
	return fmt.Sprintf(userKey, telegramID)
}
