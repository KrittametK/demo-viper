package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Configuration struct {
	App     AppConfiguration
	MongoDb DatabaseConfiguration
}
type AppConfiguration struct {
	Env   string
	Port  int
	Debug bool
}
type DatabaseConfiguration struct {
	Connection string
}

func main() {
	fmt.Println("test viper")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	var configuration Configuration
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	fmt.Println(configuration.App)
	fmt.Println(configuration.MongoDb)
}
