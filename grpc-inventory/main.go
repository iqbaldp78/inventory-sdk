package main

import (
	"fmt"
	"github.com/technical-assessment/iqbal/salestock/grpc-inventory/module"
	"github.com/technical-assessment/iqbal/salestock/grpc-inventory/module/service"
	"github.com/technical-assessment/iqbal/salestock/grpc-inventory/module/shared/lib/sqlite"
	"log"

	"github.com/spf13/viper"
)

func main() {
	setConfig("config")
	err := sqlite.New("sqlite3", viper.GetString("DB_PATH"))
	if err != nil {
		return
	}
	defer sqlite.GetDB().Close()

	log.Printf("Listening on port %v...\n", viper.GetInt("SERVER_PORT"))
	log.Fatal(module.ListenGRPC(service.New(), viper.GetInt("SERVER_PORT")))
}

func setConfig(config string) {
	viper.SetConfigName(config)
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Configuration file not found: %s", err))
	}
	viper.SetEnvPrefix("TES")
	viper.AutomaticEnv()

}
