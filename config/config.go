package config

type Config struct {
	Mysql Mysql `mapstructure:"mysql"`
	Jwt Jwt `mapstructure:"jwt"`
}
