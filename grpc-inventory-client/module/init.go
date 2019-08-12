//Package module used for prepare and manage before system run
package module

import (
	"github.com/technical-assessment/iqbal/salestock/grpc-inventory-client/module/shared/lib/sqlite"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

/*NewEngine used for create new engine*/
func NewEngine(config string) *gin.Engine {
	setConfig(config)
	router := gin.Default()
	sqlite.New("sqlite3", viper.GetString("DB_PATH"))
	setBlueprintApp(router)
	setErrorHandler(router)
	return router
}
