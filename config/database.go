package config

type DataBase struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	Options  string `yaml:"options"`
}

func (receiver DataBase) GenerateUrl() string {
	return receiver.User + ":" + receiver.Password + "@tcp(" + receiver.Host + ":" + receiver.Port + ")/" + receiver.Database + "?" + receiver.Options
}
