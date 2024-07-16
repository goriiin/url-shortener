package random

import (
	"crypto/sha256"
	"encoding/base64"
)

func GetUniqueAlias(str string) string {
	// TODO: можно переместить в конфиг
	const aliasLength = 6

	hash := sha256.Sum256([]byte(str))
	return base64.RawURLEncoding.EncodeToString(hash[:aliasLength])
}
