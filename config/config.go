package config

type Config struct {
	DataBase DataBase `yaml:"database"`
	Crontab  Crontab  `yaml:"crontab"`
}
