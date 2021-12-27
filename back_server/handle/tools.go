package handle

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

func EncodePassword(pwd string) string {
	h := sha256.New()
	h.Write([]byte(pwd))
	return hex.EncodeToString(h.Sum(nil))
}

func NewToken(username string) string {
	h := md5.New()
	h.Write([]byte(fmt.Sprintf("%s-%s", username, time.Now().String())))
	return hex.EncodeToString(h.Sum(nil))
}
