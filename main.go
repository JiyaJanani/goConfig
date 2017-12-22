package main

import (
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("No configuration file loaded - using defaults")
	}

	// If no config is found, use the default(s)
	viper.SetDefault("port", "Hello World (default)")
	theMessage := viper.GetString("port")
	fmt.Printf("\n%s\n\n", theMessage)
}
