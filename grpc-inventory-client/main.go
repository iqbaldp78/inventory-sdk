package main

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/technical-assessment/iqbal/salestock/grpc-inventory-client/module"
)

func main() {
	engine := module.NewEngine("config")
	engine.Run(fmt.Sprintf(":%v", viper.GetInt("SERVER_PORT")))
}
