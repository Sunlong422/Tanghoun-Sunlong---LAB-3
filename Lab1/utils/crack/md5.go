package crack

import (
	"crypto/md5"
	"fmt"
)

func HashMD5(text string) string {
	h := md5.Sum([]byte(text))
	return fmt.Sprintf("%x", h)
}
