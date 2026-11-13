package crack

import (
	"crypto/sha1"
	"fmt"
)

func HashSHA1(text string) string {
	h := sha1.Sum([]byte(text))
	return fmt.Sprintf("%x", h)
}
