package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	Debug    bool   `env:"APP_DEBUG,required"`
	RootPath string `env:"APP_ROOT_PATH,required"`
}

type PGConfig struct {
	ConnString       string `env:"PG_CONNECTION_STRING,required"`
	MaxPoolSize      int    `env:"PG_MAX_POOL_SIZE,required"`
	MigrationEnabled bool   `env:"PG_MIGRATION_ENABLED" envDefault:"false"`
	MigrationPath    string `env:"PG_MIGRATION_PATH" envDefault:"file://../../migrations"`
}

type Docs struct {
	Enabled  bool   `env:"SWAGGER_ENABLED" envDefault:"true"`
	Login    string `env:"SWAGGER_LOGIN" envDefault:"root"`
	Password string `env:"SWAGGER_PASSWORD" envDefault:"root"`
}

type AuthConfig struct {
	ATSecret    string `env:"AUTH_ACCESS_TOKEN_SECRET,required"`
	ATLifeMin   int    `env:"AUTH_ACCESS_TOKEN_LIFETIME_MIN,required"`
	ATName      string `env:"AUTH_ACCESS_TOKEN_NAME" envDefault:"at"`
	RTSecret    string `env:"AUTH_REFRESH_TOKEN_SECRET,required"`
	RTLifeMin   int    `env:"AUTH_REFRESH_TOKEN_LIFETIME_MIN,required"`
	RTName      string `env:"AUTH_REFRESH_TOKEN_NAME" envDefault:"rt"`
	TokenIssuer string `env:"AUTH_TOKEN_ISSUER,required"`
}

type SMTP struct {
	Host     string `env:"SMTP_HOST,required"`
	Port     int    `env:"SMTP_PORT,required"`
	User     string `env:"SMTP_USER,required"`
	Password string `env:"SMTP_PASSWORD,required"`
	Sender   string `env:"SMTP_SENDER"`
}

type Frontend struct {
	URL              string `env:"FRONTEND_URL,required"`
	ConfrimURL       string `env:"FRONTEND_CONFIRM_URL,required"`
	ResetPasswordURL string `env:"FRONTEND_RESET_PASSWORD_URL,required"`
}
type Config struct {
	App      AppConfig
	PG       PGConfig
	Auth     AuthConfig
	Docs     Docs
	SMTP     SMTP
	Frontend Frontend
}

func New() (*Config, error) {
	_ = godotenv.Load("../../.env") // If file not found try load anyway
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
