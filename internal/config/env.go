package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	MongoURL               string        `mapstructure:"MONGODB_URI"`
	DbName                 string        `mapstructure:"DB_NAME"`
	Port                   string        `mapstructure:"PORT"`
	UserCollection         string        `mapstructure:"USER_COLLECTION"`
	GridFSCollection       string        `mapstructure:"GRIDFS_COLLECTION"`
	SharedFilesCollection  string        `mapstructure:"SHARED_COLLECTION"`
	AccessTokenPrivateKey  string        `mapstructure:"ACCESS_TOKEN_PRIVATE_KEY"`
	AccessTokenPublicKey   string        `mapstructure:"ACCESS_TOKEN_PUBLIC_KEY"`
	RefreshTokenPrivateKey string        `mapstructure:"REFRESH_TOKEN_PRIVATE_KEY"`
	RefreshTokenPublicKey  string        `mapstructure:"REFRESH_TOKEN_PUBLIC_KEY"`
	AccessTokenExpiresIn   time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRED_IN"`
	RefreshTokenExpiresIn  time.Duration `mapstructure:"REFRESH_TOKEN_EXPIRED_IN"`
	AccessTokenMaxAge      int           `mapstructure:"ACCESS_TOKEN_MAXAGE"`
	RefreshTokenMaxAge     int           `mapstructure:"REFRESH_TOKEN_MAXAGE"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	err = viper.Unmarshal(&config)

	if err != nil {
		return Config{}, err
	}

	return config, nil
}
