package crud

import (
	"fmt"
	"time"

	"github.com/gocraft/dbr"

	"github.com/technical-assessment/iqbal/salestock/grpc-inventory/module/shared/object"
)

//ReportSalesHeader used for store header report
type ReportSalesHeader struct {
	PrintDate        string
	ParamDate        string
	TotalOmset       float32
	TotalLabaGross   float32
	TotalPenjualan   int64
	TotalBarang      int64
	ReportSalesLines []ReportSalesLines
}

//ReportSalesLines used for store lines report
type ReportSalesLines struct {
	OrderNum      string  `db:"order_num" json:"order_num"`
	OrderDate     string  `db:"order_date" json:"order_date"`
	Sku           string  `db:"sku" json:"sku"`
	Itemname      string  `db:"item_name" json:"item_name"`
	QtySo         int64   `db:"qty_so" json:"qty_so"`
	SellingPrice  float32 `db:"selling_price" json:"selling_price"`
	TotalPrice    float32 `db:"total_price" json:"total_price"`
	PurchasePrice float32 `db:"purchase_price" json:"purchase_price"`
	Laba          float32 `db:"laba" json:"laba"`
}

//SearchReportSales used for interact to DB. to get sales data
func (obj *ReportSalesHeader) SearchReportSales(db *dbr.Session, fromDate, toDate string) (output object.SearchResult) {
	tx, err := db.Begin()
	defer tx.RollbackUnlessCommitted()
	if err != nil {
		return
	}
	var dataLines []ReportSalesLines
	var dataHeader []ReportSalesHeader
	var totalOmset, labaKotor float32
	var totalSales, totalBarang int64
	var condition string

	if fromDate != "" && toDate != "" {
		condition = fmt.Sprintf("WHERE order_date between '%v' and '%v'", fromDate, toDate)
	}

	queryLines := fmt.Sprintf(`SELECT so.order_num, so.order_date,sol.sku,inv.item_name,sol.qty_so,sol.selling_price,sol.selling_price*sol.qty_so total_price, pol.purchase_price,((sol.selling_price*sol.qty_so)-(pol.purchase_price*sol.qty_so))laba from sales_order so join so_lines sol on so.order_num = sol.order_num join inventory inv on sol.sku = inv.sku join po_lines pol on so.po_line_id=pol.po_line_id %v`, condition)
	_, err = tx.SelectBySql(queryLines).Load(&dataLines)

	querySumSales := fmt.Sprintf(`select count(*)as cntSales from shipping where status ='completed'`)
	err = tx.SelectBySql(querySumSales).LoadOne(&totalSales)

	if err == nil {
		tx.Commit()
	}
	for i := range dataLines {
		totalOmset = totalOmset + dataLines[i].TotalPrice
		labaKotor = labaKotor + dataLines[i].Laba
		totalBarang = totalBarang + dataLines[i].QtySo
	}
	header := ReportSalesHeader{
		PrintDate:        time.Now().Format("02 January 2006"),
		ParamDate:        fromDate + " s/d " + toDate,
		TotalOmset:       totalOmset,
		TotalLabaGross:   labaKotor,
		TotalPenjualan:   totalSales,
		TotalBarang:      totalBarang,
		ReportSalesLines: dataLines,
	}
	dataHeader = append(dataHeader, header)
	output.Data = dataHeader
	return output
}
