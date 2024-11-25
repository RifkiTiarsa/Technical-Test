package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	Driver   string
}

type ApiConfig struct {
	ApiPort string
}

type TokenConfig struct {
	IssuerName       string `json:"IssuerName"`
	JwtSignatureKy   []byte `json:"JwtSignatureKy"`
	JwtSigningMethod *jwt.SigningMethodHMAC
	JwtExpiresTime   time.Duration
}

type Config struct {
	DBConfig
	ApiConfig
	TokenConfig
}

func (c *Config) readConfig() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("error loading.env file")
	}
	// Read DBConfig from environment variables, file, etc.
	c.DBConfig = DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		Driver:   os.Getenv("DB_DRIVER"),
	}

	// Read ApiConfig from environment variables, file, etc.
	c.ApiConfig = ApiConfig{
		ApiPort: os.Getenv("API_PORT"),
	}

	// Read TokenConfig from environment variables, file, etc.
	tokenExpire, _ := strconv.Atoi(os.Getenv("TOKEN_EXPIRE"))
	c.TokenConfig = TokenConfig{
		IssuerName:       os.Getenv("TOKEN_ISSUE"),
		JwtSignatureKy:   []byte(os.Getenv("TOKEN_SECRET")),
		JwtSigningMethod: jwt.SigningMethodHS256,
		JwtExpiresTime:   time.Duration(tokenExpire) * time.Minute,
	}

	if c.Host == "" || c.Port == "" || c.Username == "" || c.Password == "" || c.DBName == "" || c.Driver == "" || c.ApiPort == "" || c.JwtExpiresTime < 0 || len(c.JwtSignatureKy) == 0 {
		return fmt.Errorf("missing required environment variables")
	}

	return nil
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := cfg.readConfig(); err != nil {
		return nil, err
	}

	return cfg, nil
}
