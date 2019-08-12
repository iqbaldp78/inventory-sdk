/*Package inventory used for inventory module*/
package inventory

import (
	"fmt"
	"net/http"

	"github.com/satori/go.uuid"

	"github.com/technical-assessment/iqbal/salestock/grpc-inventory-client/module/csv"
	"github.com/technical-assessment/iqbal/salestock/grpc-inventory-client/module/shared/client"
	pb "github.com/technical-assessment/iqbal/salestock/grpc-inventory-client/module/shared/client/pb"
	"github.com/technical-assessment/iqbal/salestock/grpc-inventory-client/module/shared/object"

	"github.com/gin-gonic/gin"
)

//SearchOnhanQTY used as controller for search onhandqty
func SearchOnhanQTY(client *client.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := client.GetOnhandQTY(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		csv := csv.NewCsv(data, nil, nil, nil, nil)
		u2 := uuid.NewV4()
		fileName := fmt.Sprintf("catatan_jumlah_barang_%v.csv", u2)
		err = csv.GenerateCSVOnhand(fileName)
		fmt.Println(err)
		header := c.Writer.Header()
		header["Content-type"] = []string{"application/octet-stream"}
		header["Content-Disposition"] = []string{"attachment; filename= " + fileName}
		c.File(fileName)
		csv.DeleteFile(fileName)
	}
}

//SearchInboundGoods used as controller for search InboundGoods
func SearchInboundGoods(client *client.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := client.GetInbounds(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		// c.JSON(http.StatusOK, data)
		csv := csv.NewCsv(nil, data, nil, nil, nil)
		u2 := uuid.NewV4()
		fileName := fmt.Sprintf("catatan_barang_masuk_%v.csv", u2)
		err = csv.GenerateInbound(fileName)
		fmt.Println(err)
		header := c.Writer.Header()
		header["Content-type"] = []string{"application/octet-stream"}
		header["Content-Disposition"] = []string{"attachment; filename= " + fileName}
		c.File(fileName)
		csv.DeleteFile(fileName)
	}
}

//SearchOutboundGoods used as controller for search SearchOutboundGoods
func SearchOutboundGoods(client *client.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := client.GetOutbounds(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		// c.JSON(http.StatusOK, data)
		csv := csv.NewCsv(nil, nil, data, nil, nil)
		u2 := uuid.NewV4()
		fileName := fmt.Sprintf("catatan_barang_keluar_%v.csv", u2)
		err = csv.GenerateOutbound(fileName)
		fmt.Println(err)
		header := c.Writer.Header()
		header["Content-type"] = []string{"application/octet-stream"}
		header["Content-Disposition"] = []string{"attachment; filename= " + fileName}
		c.File(fileName)
		csv.DeleteFile(fileName)
	}

}

//SearchReportValue used as controller for search SearchReportValue
func SearchReportValue(client *client.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := client.GetReportValue(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		// c.JSON(http.StatusOK, data)
		csv := csv.NewCsv(nil, nil, nil, data, nil)
		u2 := uuid.NewV4()
		fileName := fmt.Sprintf("laporan_nilai_barang_%v.csv", u2)
		err = csv.GenerateReportValue(fileName)
		fmt.Println(err)
		header := c.Writer.Header()
		header["Content-type"] = []string{"application/octet-stream"}
		header["Content-Disposition"] = []string{"attachment; filename= " + fileName}
		c.File(fileName)
		csv.DeleteFile(fileName)
	}
}

//SearchReportSales used as controller for search SearchReportSales
func SearchReportSales(client *client.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var param object.SearchParamDate
		if err := c.ShouldBindQuery(&param); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		paramQuery := pb.ParamDate{
			DateFrom: param.DateFrom,
			DateTo:   param.DateTo,
		}
		data, err := client.GetReportSales(c, &paramQuery)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		csv := csv.NewCsv(nil, nil, nil, nil, data)
		u2 := uuid.NewV4()
		fileName := fmt.Sprintf("laporan_nilai_barang_%v.csv", u2)
		err = csv.GenerateReportSales(fileName)
		fmt.Println(err)
		header := c.Writer.Header()
		header["Content-type"] = []string{"application/octet-stream"}
		header["Content-Disposition"] = []string{"attachment; filename= " + fileName}
		c.File(fileName)
		csv.DeleteFile(fileName)
	}
}
