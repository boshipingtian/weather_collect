package config

type Crontab struct {
	Open  bool   `yaml:"open"`
	Works []Work `yaml:"works"`
}

type Work struct {
	Name    string `yaml:"name"`
	Crontab string `yaml:"crontab"`
}
