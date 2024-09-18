package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Database []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
	} `mapstructure:"database"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./config/") // path to config
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	// read configuration
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fail to read configuration %w", err))
	}
	// read server configuration
	fmt.Println("Server Port::", viper.GetInt("server.port"))
	fmt.Println("Sercurity key::", viper.GetString("sercurity.jwt.key"))

	// configure structrure
	var config Config
	if err = viper.Unmarshal(&config); err != nil {
		fmt.Printf("unable to decode configuration %v", err)
	}
	fmt.Println("Config Port::", config.Server.Port)

	for _, db := range config.Database {
		fmt.Printf("Database User: %s, Password: %s, Host: %s \n", db.User, db.Password, db.Host)
	}
}
