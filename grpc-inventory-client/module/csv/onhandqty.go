package csv

import (
	"encoding/csv"
	// "fmt"
	"log"
	"strconv"

	pb "github.com/technical-assessment/iqbal/salestock/grpc-inventory-client/module/shared/client/pb"
	"os"
)

//NewCsv create new csv struct
func NewCsv(dataOnhand *pb.ResOnhandQTY, dataInbound *pb.ResInbounds, dataOutbound *pb.ResOutbounds, dataRepValue *pb.ResReportValue, dataRepSales *pb.ResReportSalesHeader) ClientCsv {
	return ClientCsv{
		dataOnhandQty: dataOnhand,
		dataInbound:   dataInbound,
		dataOutbound:  dataOutbound,
		dataRepValue:  dataRepValue,
		dataRepSales:  dataRepSales,
	}
}

//ClientCsv --
type ClientCsv struct {
	template      [][]string
	dataOnhandQty *pb.ResOnhandQTY
	dataInbound   *pb.ResInbounds
	dataOutbound  *pb.ResOutbounds
	dataRepValue  *pb.ResReportValue
	dataRepSales  *pb.ResReportSalesHeader
}

//GenerateCSVOnhand --
func (c *ClientCsv) GenerateCSVOnhand(fileName string) error {
	file, err := c.createFile(fileName)
	if err != nil {
		return err
	}

	c.template = [][]string{
		{"SKU", ";", "Nama Item", ";", "Jumlah Sekarang"},
	}

	res := c.appendData("onhand")
	c.template = append(c.template, res...)
	log.Println(c.template[1])

	err = c.writeCsvtemplate(file)
	if err != nil {
		return err
	}
	defer file.Close()
	return err
}

func (c *ClientCsv) appendData(typed string) [][]string {
	var data2 [][]string
	if typed == "onhand" {
		for i := range c.dataOnhandQty.Data {
			var data []string
			stock := strconv.Itoa(int(c.dataOnhandQty.Data[i].Stock))
			data = append(data, c.dataOnhandQty.Data[i].Sku, ";", c.dataOnhandQty.Data[i].Itemname, ";", stock, ";")
			data2 = append(data2, data)
		}
	} else if typed == "inbound" {
		for i := range c.dataInbound.Data {
			var data []string
			qtyPo := strconv.Itoa(int(c.dataInbound.Data[i].QtyPo))
			qtyRcv := strconv.Itoa(int(c.dataInbound.Data[i].QtyReceive))
			// hargaBeli := fmt.Sprintf("%f", c.dataInbound.Data[i].PurchasePrice)
			hargaBeli := strconv.FormatFloat(float64(c.dataInbound.Data[i].PurchasePrice), 'f', 5, 64)
			total := strconv.FormatFloat(float64(c.dataInbound.Data[i].TotalInboundPrice), 'f', 5, 64)
			data = append(data, c.dataInbound.Data[i].PoDate, ";", c.dataInbound.Data[i].Sku, ";", c.dataInbound.Data[i].ItemName, ";", qtyPo, ";", qtyRcv, ";")
			data = append(data, hargaBeli, ";", total, ";", c.dataInbound.Data[i].KwitansiNum, ";", c.dataInbound.Data[i].Noted)
			data2 = append(data2, data)
		}
	} else if typed == "outbound" {
		for i := range c.dataOutbound.Data {
			var data []string
			jmlKeluar := strconv.Itoa(int(c.dataOutbound.Data[i].QtyOut))
			if c.dataOutbound.Data[i].SellingPrice == -1 {
				c.dataOutbound.Data[i].SellingPrice = 0
			}
			if c.dataOutbound.Data[i].TotalOutboundPrice == -1 {
				c.dataOutbound.Data[i].TotalOutboundPrice = 0
			}
			total := strconv.FormatFloat(float64(c.dataOutbound.Data[i].TotalOutboundPrice), 'f', 5, 64)
			hargaJual := strconv.FormatFloat(float64(c.dataOutbound.Data[i].SellingPrice), 'f', 5, 64)
			data = append(data, c.dataOutbound.Data[i].OrderDate, ";", c.dataOutbound.Data[i].Sku, ";", c.dataOutbound.Data[i].Itemname, ";", jmlKeluar, ";", hargaJual, ";", total, ";", c.dataOutbound.Data[i].Noted)
			data2 = append(data2, data)
		}
	}
	if typed == "repValue" {
		for i := range c.dataRepValue.Data[0].ReportValueLines {
			var data []string
			jumlah := strconv.Itoa(int(c.dataRepValue.Data[0].ReportValueLines[i].QtyTotalInv))
			rata2 := strconv.FormatFloat(float64(c.dataRepValue.Data[0].ReportValueLines[i].AvgPruchasePrice), 'f', 5, 64)
			total := strconv.FormatFloat(float64(c.dataRepValue.Data[0].ReportValueLines[i].TotalPurchasePrice), 'f', 5, 64)
			data = append(data, c.dataRepValue.Data[0].ReportValueLines[i].Sku, ";", c.dataRepValue.Data[0].ReportValueLines[i].Itemname, ";", jumlah, ";", rata2, ";", total)
			data2 = append(data2, data)
		}

	}
	if typed == "repSales" {
		for i := range c.dataRepSales.Data[0].ReportSalesLines {
			var data []string
			jumlah := strconv.Itoa(int(c.dataRepSales.Data[0].ReportSalesLines[0].QtySo))
			hargaJual := strconv.FormatFloat(float64(c.dataRepSales.Data[0].ReportSalesLines[i].SellingPrice), 'f', 5, 64)
			total := strconv.FormatFloat(float64(c.dataRepSales.Data[0].ReportSalesLines[i].TotalPrice), 'f', 5, 64)
			hargaBeli := strconv.FormatFloat(float64(c.dataRepSales.Data[0].ReportSalesLines[i].PurchasePrice), 'f', 5, 64)
			laba := strconv.FormatFloat(float64(c.dataRepSales.Data[0].ReportSalesLines[i].Laba), 'f', 5, 64)
			data = append(data, c.dataRepSales.Data[0].ReportSalesLines[0].OrderNum, ";", c.dataRepSales.Data[0].ReportSalesLines[0].OrderDate, ";", c.dataRepSales.Data[0].ReportSalesLines[0].Sku, ";", c.dataRepSales.Data[0].ReportSalesLines[0].ItemName)
			data = append(data, ";", jumlah, ";", hargaJual, ";", total, ";", hargaBeli, ";", laba)
			data2 = append(data2, data)
		}

	}

	return data2
}

func (c *ClientCsv) createFile(fileName string) (*os.File, error) {
	file, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}
	return file, err
}

func (c *ClientCsv) writeCsvHeader(file *os.File, value [][]string) error {
	w := csv.NewWriter(file)
	w.Comma = '\t'
	for _, value := range c.template {
		if err := w.Write(value); err != nil {
			return err
		}
	}
	defer w.Flush()
	return nil
}

func (c *ClientCsv) writeCsvtemplate(file *os.File) error {
	w := csv.NewWriter(file)
	w.Comma = '\t'
	for _, value := range c.template {
		if err := w.Write(value); err != nil {
			return err
		}
	}
	defer w.Flush()
	return nil
}

//DeleteFile --
func (c *ClientCsv) DeleteFile(path string) error {
	// delete file
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return err
}
