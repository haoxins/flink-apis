package rest

import (
	"crypto/md5"
	"encoding/hex"
)

func GenJobId(namespace, name string) string {
	hash := md5.Sum([]byte(namespace + name))
	return hex.EncodeToString(hash[:])
}
