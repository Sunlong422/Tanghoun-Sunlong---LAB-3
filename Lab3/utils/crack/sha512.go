package crack

import (
	"crypto/sha512"
	"fmt"
)

func HashSHA512(text string) string {
	h := sha512.Sum512([]byte(text))
	return fmt.Sprintf("%x", h)
}
