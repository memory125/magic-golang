package conf

// MySQLConf ...
type MySQLConf struct {
	IP       string `ini:"ip"`
	Port     string `ini:"port"`
	User     string `ini:"user"`
	Password string `ini:"password"`
	Database string `ini:"database"`
}

// RedisConf ...
type RedisConf struct {
	IP   string `ini:"ip"`
	Port string `ini:"port"`
}

type ProjectConf struct {
	MySQLConf `ini:"mysql"`
	RedisConf `ini:"redis"`
}
