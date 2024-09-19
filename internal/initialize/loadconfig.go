package initialize

import (
	"fmt"

	"github.com/danhbuidcn/go_backend_api/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
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
	if err = viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("unable to decode configuration %v", err)
	}
}
