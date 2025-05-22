package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	Debug                  bool   `env:"APP_DEBUG,required"`
	RootPath               string `env:"APP_ROOT_PATH,required"`
	AccVerificationEnabled bool   `env:"APP_ACC_VERIFICATION_ENABLED" envDefault:"true"`
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
type Config struct {
	App  AppConfig
	PG   PGConfig
	Auth AuthConfig
	Docs Docs
	SMTP SMTP
}

type SMTP struct {
	Host     string `env:"SMTP_HOST"`
	Port     int    `env:"SMTP_PORT"`
	User     string `env:"SMTP_USER"`
	Password string `env:"SMTP_PASSWORD"`
	Sender   string `env:"SMTP_SENDER"` // may be unset when, sender is User
}

func New() (*Config, error) {
	_ = godotenv.Load("../../.env") // If file not found try load anyway
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	if cfg.App.AccVerificationEnabled {
		missingSMTP := cfg.SMTP.Host == "" || cfg.SMTP.Port == 0 || cfg.SMTP.User == "" || cfg.SMTP.Password == ""
		if missingSMTP {
			return nil, fmt.Errorf("SMTP configuration is required when account verification is enabled. Either disable APP_ACC_VERIFICATION_ENABLED or set all required SMTP environment variables")
		}
	}
	return cfg, nil
}
