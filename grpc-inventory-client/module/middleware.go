package module

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/technical-assessment/iqbal/salestock/grpc-inventory-client/module/shared/message"
)

//setConfig used for setup application configuration
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

/*setErrorHandler used for handling 404 request*/
func setErrorHandler(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		msg := message.New(0, nil)
		c.AbortWithStatusJSON(msg.StatusCode, msg)
	})
}
