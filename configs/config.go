package configs

type ServerConfig struct {
	Name         string      `mapstructure:"name"`
	Mode         string      `mapstructure:"mode"`
	TimeLocation string      `mapstructure:"time_location"`
	MySQLConfig  *MySQLConf  `mapstructure:"mysql"`
	ConsulConfig *ConsulConf `mapstructure:"consul"`
}

type MySQLConf struct {
	DataSourceName  string `mapstructure:"dsn"`
	MaxIdleConn     int    `mapstructure:"max_idle_conn"`
	MaxOpenConn     int    `mapstructure:"max_open_conn"`
	MaxConnLifeTime int    `mapstructure:"max_conn_life_time"`
}

type ConsulConf struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}
