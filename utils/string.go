package utils

import (
	"fmt"
	"crypto/md5"
)

func Md5(buf []byte) string {
    hash := md5.New()
    hash.Write(buf)
    return fmt.Sprintf("%x", hash.Sum(nil))
}