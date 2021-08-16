package models

type Task struct {
	Cron   string
	Path   string
	Enable bool
	Mode   string //obo alo
	Word   string
}

type Env struct {
	Name  string
	Value string
}
