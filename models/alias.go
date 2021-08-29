package models

type Alias struct {
	ID    int
	Key   string
	Value string
}

// alias /asset $1 => run jd_bean_change.js $1 -w
