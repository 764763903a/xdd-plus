package models

type Env struct {
	ID    int
	Name  string `gorm:"unique"`
	Value string
}

func ExportEnv(env *Env) {
	value := env.Value
	if err := db.Where("name = ?", env.Name).First(env).Error; err != nil {
		db.Create(env)
	} else {
		db.Model(env).Update("value", value)
	}
}

func UnExportEnv(env *Env) {
	db.Where("name = ?", env.Name).Delete(env)
}

func GetEnvs() []Env {
	envs := []Env{}
	db.Find(&envs)
	return envs
}

func GetEnv(name string) string {
	env := &Env{}
	db.Where("name = ?", name).First(env)
	return env.Value
}
