package models

import (
	"time"
)

type Token struct {
	Expiration time.Time
	Token      string
	Address    string
}
