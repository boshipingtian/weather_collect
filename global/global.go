package global

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"weather-colly/config"
)

var (
	Config *config.Config
	DB     *gorm.DB
	Logger *logrus.Logger
)
