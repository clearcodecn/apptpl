package server

type Option func(config *Config)

type Config struct {
	MysqlDSN          string
	EnableMockStorage bool
	Addr string
}

// root:root@tcp(localhost:33060)/clearcode?charset=utf8
func WithMysqlDSN(dsn string) Option {
	return func(c *Config) {
		c.MysqlDSN = dsn
	}
}

func EnableMockStorage() Option {
	return func(c *Config) {
		c.EnableMockStorage = true
	}
}

func WithAddr(addr string) Option {
	return func(config *Config) {
		config.Addr = addr
	}
}