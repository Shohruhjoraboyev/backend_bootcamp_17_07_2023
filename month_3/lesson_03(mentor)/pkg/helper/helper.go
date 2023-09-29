package helper

import (
	"strconv"
	"strings"

	"github.com/xtgo/uuid"
	"golang.org/x/crypto/bcrypt"
)

func ReplaceQueryParams(namedQuery string, params map[string]interface{}) (string, []interface{}) {
	var (
		i    int = 1
		args []interface{}
	)

	for k, v := range params {
		if k != "" && strings.Contains(namedQuery, ":"+k) {
			namedQuery = strings.ReplaceAll(namedQuery, ":"+k, "$"+strconv.Itoa(i))
			args = append(args, v)
			i++
		}
	}

	return namedQuery, args
}
func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

func GeneratePasswordHash(pass string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pass), 10)
}
func ComparePasswords(hashedPass, pass []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPass, pass)
}
