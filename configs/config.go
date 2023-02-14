package configs

type ServerConfig struct {
	Name         string      `mapstructure:"name"`
	Mode         string      `mapstructure:"mode"`
	TimeLocation string      `mapstructure:"time_location"`
	ConfigPath   string      `mapstructure:"config_path"`
	MySQLConfig  *MySQLConf  `mapstructure:"mysql"`
	ConsulConfig *ConsulConf `mapstructure:"consul"`
	MinioConfig  *MinioConf  `mapstructure:"minio"`
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

type MinioConf struct {
	EndPoint        string `mapstructure:"end_point"`
	AccessKeyId     string `mapstructure:"access_key_id"`
	SecretAccessKey string `mapstructure:"secret_access_key"`
	UserSsl         bool   `mapstructure:"user_ssl"`
}
