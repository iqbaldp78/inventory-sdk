package csv

import (
	"strconv"
)

//GenerateReportSales --
func (c *ClientCsv) GenerateReportSales(fileName string) error {
	file, err := c.createFile(fileName)
	if err != nil {
		return err
	}

	TotalOmset := strconv.FormatFloat(float64(c.dataRepSales.Data[0].TotalOmset), 'f', 5, 64)
	TotalLabaGross := strconv.FormatFloat(float64(c.dataRepSales.Data[0].TotalLabaGross), 'f', 5, 64)
	TotalBarang := strconv.FormatFloat(float64(c.dataRepSales.Data[0].TotalBarang), 'f', 5, 64)
	TotalPenjualan := strconv.FormatFloat(float64(c.dataRepSales.Data[0].TotalPenjualan), 'f', 5, 64)
	c.template = [][]string{
		{"LAPORAN PENJUALAN"},
		{";"},
		{"Tanggal Cetak : ", c.dataRepSales.Data[0].PrintDate},
		{"Tanggal : ", c.dataRepSales.Data[0].ParamDate},
		{"Total Omzet : ", TotalOmset},
		{"Total Laba Kotor : ", TotalLabaGross},
		{"Total Penjualan : ", TotalPenjualan},
		{"Total Barang : ", TotalBarang},
		{";"},
		{"ID Pesanan", ";", "Waktu", ";", "SKU", ";", "Nama Barang", ";", "Jumlah", ";", "Harga Jual", ";", "Total", ";", "Harga Beli", ";", "Laba"},
	}

	res := c.appendData("repSales")
	c.template = append(c.template, res...)

	err = c.writeCsvtemplate(file)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}
