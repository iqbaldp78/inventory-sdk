package csv

//GenerateOutbound --
func (c *ClientCsv) GenerateOutbound(fileName string) error {
	file, err := c.createFile(fileName)
	if err != nil {
		return err
	}
	c.template = [][]string{
		{"Waktu", ";", "SKU", ";", "Nama Barang", ";", "Jumlah Keluar", ";", "Harga Jual", ";", "Total", ";", "Catatan"},
	}

	res := c.appendData("outbound")
	c.template = append(c.template, res...)

	err = c.writeCsvtemplate(file)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}
