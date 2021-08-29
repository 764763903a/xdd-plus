package models

import (
	"time"
)

type Wish struct {
	ID         int
	CreatedAt  time.Time
	UserNumber int
	Content    string
	Coin       int
	Status     int // 1 2
}
