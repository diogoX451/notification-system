package gorm

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ConfigGorm struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
	Debug    bool
	Charset  string
	Loc      *time.Location
	Driver   string
	Migrate  []interface{}
}

func NewConfigGorm(cfg ...func(*ConfigGorm)) *ConfigGorm {
	svr := &ConfigGorm{}
	for _, f := range cfg {
		f(svr)
	}

	return svr
}

func (c *ConfigGorm) getDSNPostgres() string {
	return "host=" + c.Host + " port=" + c.Port + " user=" + c.User + " dbname=" + c.DBName + " password=" + c.Password + " sslmode=" + c.SSLMode
}

func (c *ConfigGorm) Connect() (*gorm.DB, error) {
	switch c.Driver {
	case "postgres":
		return c.connectPostgres()
	default:
		return nil, nil
	}
}

func (c *ConfigGorm) connectPostgres() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(c.getDSNPostgres()), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(c.Migrate...); err != nil {
		return nil, err
	}

	return db, nil
}

func WithHost(host string) func(*ConfigGorm) {
	return func(c *ConfigGorm) {
		c.Host = host
	}
}

func WithPort(port string) func(*ConfigGorm) {
	return func(c *ConfigGorm) {
		c.Port = port
	}
}

func WithUser(user string) func(*ConfigGorm) {
	return func(c *ConfigGorm) {
		c.User = user
	}
}

func WithPassword(password string) func(*ConfigGorm) {
	return func(c *ConfigGorm) {
		c.Password = password
	}
}

func WithDBName(dbname string) func(*ConfigGorm) {
	return func(c *ConfigGorm) {
		c.DBName = dbname
	}
}

func WithSSLMode(sslmode string) func(*ConfigGorm) {
	return func(c *ConfigGorm) {
		c.SSLMode = sslmode
	}
}

func WithDebug(debug bool) func(*ConfigGorm) {
	return func(c *ConfigGorm) {
		c.Debug = debug
	}
}

func WithCharset(charset string) func(*ConfigGorm) {
	return func(c *ConfigGorm) {
		c.Charset = charset
	}
}

func WithLoc(loc *time.Location) func(*ConfigGorm) {
	return func(c *ConfigGorm) {
		c.Loc = loc
	}
}

func WithDriver(driver string) func(*ConfigGorm) {
	return func(c *ConfigGorm) {
		c.Driver = driver
	}
}

func WithMigrate(migrate []interface{}) func(*ConfigGorm) {
	return func(c *ConfigGorm) {
		c.Migrate = migrate
	}
}
