package module

import (
	"net/http"

	"github.com/technical-assessment/iqbal/salestock/grpc-inventory-client/module/inventory"
	"github.com/technical-assessment/iqbal/salestock/grpc-inventory-client/module/shared/client"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func setBlueprintApp(router *gin.Engine) {
	client, _ := client.NewClient(viper.GetString("INVENTORY_SERVER"))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	web := router.Group("web")
	{
		web.GET("/onhandqty", inventory.SearchOnhanQTY(client))
		web.GET("/inboundgoods", inventory.SearchInboundGoods(client))
		web.GET("/outboundsgoods", inventory.SearchOutboundGoods(client))
		web.GET("/reportvalue", inventory.SearchReportValue(client))
		web.GET("/reportsales", inventory.SearchReportSales(client))
	}

}
