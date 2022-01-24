package conf

// Yaml struct of conf
type Yaml struct {
	Mysql struct {
		User     string `conf:"user"`
		Host     string `conf:"host"`
		Password string `conf:"password"`
		Port     string `conf:"port"`
		Name     string `conf:"name"`
	}
	Cache struct {
		Enable bool     `conf:"enable"`
		List   []string `conf:"list,flow"`
	}
}

// Yaml1 struct of conf
type Yaml1 struct {
	SQLConf   Mysql `conf:"mysql"`
	CacheConf Cache `conf:"cache"`
}

// Yaml2 struct of conf
type Yaml2 struct {
	Mysql `conf:"mysql,inline"`
	Cache `conf:"cache,inline"`
}

// Mysql struct of mysql conf
type Mysql struct {
	User     string `conf:"user"`
	Host     string `conf:"host"`
	Password string `conf:"password"`
	Port     string `conf:"port"`
	Name     string `conf:"name"`
}

// Cache struct of cache conf
type Cache struct {
	Enable bool     `conf:"enable"`
	List   []string `conf:"list,flow"`
}
