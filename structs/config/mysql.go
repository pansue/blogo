package config

type Mysql struct {
	User               string `mapstructure:"user"`
	Password           string `mapstructure:"password"`
	Database           string `mapstructure:"database"`
	Host               string `mapstructure:"host"`
	Charset            string `mapstructure:"charset"`
	MaxIdleConnections int    `mapstructure:"max-idle-connections"`
	MaxOpenConnections int    `mapstructure:"max-open-connections"`
}
