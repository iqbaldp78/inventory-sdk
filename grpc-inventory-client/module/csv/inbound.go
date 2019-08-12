package csv

//GenerateInbound --
func (c *ClientCsv) GenerateInbound(fileName string) error {
	file, err := c.createFile(fileName)
	if err != nil {
		return err
	}
	c.template = [][]string{
		{"Waktu", ";", "SKU", ";", "Nama Barang", ";", "Jumlah Pesanan", ";", "Jumlah Diterima", ";", "Harga Beli", ";", "Total", ";", "Nomor Kwitansi", ";", "Catatan", ";"},
	}

	res := c.appendData("inbound")
	c.template = append(c.template, res...)

	err = c.writeCsvtemplate(file)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}
