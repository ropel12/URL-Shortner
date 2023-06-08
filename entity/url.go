package entity

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type URL struct {
	LongURL  string `json:"long_url"`
	ShortURL string `json:"short_url"`
}

func GetRandomShortURL(longURL string) string {
	s := fmt.Sprintf("%s%d", longURL, time.Now().Unix())
	sum := sha256.Sum256([]byte(s))
	hashString := fmt.Sprintf("%x", sum)
	return hashString[:8]
}
