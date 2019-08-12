package crud

import (
	"fmt"
	"time"

	"github.com/gocraft/dbr"

	"github.com/technical-assessment/iqbal/salestock/grpc-inventory/module/shared/object"
)

//ReportValueHeader used for store header report
type ReportValueHeader struct {
	CountSku         int64   `db:"count_sku" json:"count_sku" binding:"required"`
	SumStock         int64   `db:"sum_stock" json:"sum_stock" binding:"required"`
	TotalValue       float32 `db:"-" json:"total_value" binding:"required"`
	PrintDate        string  `db:"-" json:"print_date,omitempty"`
	ReportValueLines []ReportValueLines
}

//ReportValueLines used for store lines report
type ReportValueLines struct {
	Sku                string  `db:"sku" json:"sku" binding:"required"`
	Itemname           string  `db:"item_name" json:"item_name" binding:"required"`
	QtyTotalInv        int64   `db:"qty_total_inv" json:"qty_total_inv" binding:"required"`
	AvgPruchasePrice   float32 `db:"avg_purchase_price" json:"avg_purchase_price" binding:"required"`
	TotalPurchasePrice float32 `db:"-" json:"total_purchase_price" binding:"required"`
}

//SearchReportValue used for interact to DB. to get outbound data
func (obj *ReportValueHeader) SearchReportValue(db *dbr.Session) (output object.SearchResult) {
	tx, err := db.Begin()
	defer tx.RollbackUnlessCommitted()
	if err != nil {
		return
	}
	var dataLines []ReportValueLines
	var dataHeader []ReportValueHeader
	var TotalNilai float32

	total := fmt.Sprintf(`select COUNT(*) from inventory inv join po_lines pol on inv.sku = pol.sku`)
	err = tx.SelectBySql(total).LoadOne(&output)

	queryHeader := fmt.Sprintf(`select COUNT(*)count_sku,sum(stock)sum_stock from inventory `)
	err = tx.SelectBySql(queryHeader).LoadOne(&dataHeader)

	queryLines := fmt.Sprintf(`select inv.sku,inv.item_name,inv.stock qty_total_inv,AVG(pol.purchase_price) avg_purchase_price from inventory inv join po_lines pol on inv.sku = pol.sku group by inv.sku,inv.item_name,inv.stock`)
	_, err = tx.SelectBySql(queryLines).Load(&dataLines)
	if err == nil {
		tx.Commit()
	}
	for i := range dataLines {
		dataLines[i].TotalPurchasePrice = dataLines[i].AvgPruchasePrice * float32(dataLines[i].QtyTotalInv)
		TotalNilai = TotalNilai + dataLines[i].TotalPurchasePrice
	}
	for i := range dataHeader {
		dataHeader[i].PrintDate = time.Now().Format("02 January 2006")
		dataHeader[i].TotalValue = TotalNilai
		dataHeader[i].ReportValueLines = dataLines
	}
	output.Data = dataHeader
	return output
}
