package models

import (
	"time"
)

type Token struct {
	ID        int
	CreatedAt time.Time
	Token     string
}
