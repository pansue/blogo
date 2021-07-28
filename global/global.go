package global

import (
	"blogo/config"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GORM  *gorm.DB
	VIPER *viper.Viper
	CONF  config.Config
	LOG   *zap.Logger
)
