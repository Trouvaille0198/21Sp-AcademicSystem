package config

// ServerConfig 服务器配置项
type ServerConfig struct {
	ServerName string     `mapstructure:"server-name"`
	Port       int        `mapstructure:"port"`
	GormInfo   GormConfig `mapstructure:"gorm-info"`
}

// GormConfig gorm连接数据库的配置项
type GormConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbName"`
}
