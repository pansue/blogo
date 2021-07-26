package setup

import (
	"blogo/global"
	"blogo/utils"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func Viper() *viper.Viper {
	v := viper.New()
	v.SetConfigFile(utils.CONFIGPATH)
	v.SetConfigType("toml")

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error in reading config file: %v", err))
	}

	if err := v.Unmarshal(&global.CONF); err != nil {
		fmt.Println(err)
	}

	global.CONF.Mysql.Password = os.ExpandEnv(global.CONF.Mysql.Password)

	return v
}
