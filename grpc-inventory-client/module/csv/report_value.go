package csv

import (
	"strconv"
)

//GenerateReportValue --
func (c *ClientCsv) GenerateReportValue(fileName string) error {
	file, err := c.createFile(fileName)
	if err != nil {
		return err
	}

	CountSku := strconv.Itoa(int(c.dataRepValue.Data[0].CountSku))
	SumStock := strconv.Itoa(int(c.dataRepValue.Data[0].SumStock))
	TotalValue := strconv.FormatFloat(float64(c.dataRepValue.Data[0].TotalValue), 'f', 5, 64)
	c.template = [][]string{
		{"LAPORAN NILAI BARANG"},
		{";"},
		{"Tanggal Cetak : ", c.dataRepValue.Data[0].PrintDate},
		{"Jumlah SKU : ", CountSku},
		{"Jumlah Total Barang : ", SumStock},
		{"Total Nilai : ", TotalValue},
		{";"},
		{"SKU", ";", "Nama Item", ";", "Jumlah", ";", "Rata-Rata Harga Beli", ";", "Total"},
	}

	res := c.appendData("repValue")
	c.template = append(c.template, res...)

	err = c.writeCsvtemplate(file)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}
