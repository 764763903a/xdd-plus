package models

import (
	"time"
)

type Token struct {
	expiration time.Time
	Token      string
}
