package crud

import (
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/technical-assessment/iqbal/salestock/grpc-inventory/module/shared/object"
	"github.com/technical-assessment/iqbal/salestock/grpc-inventory/module/shared/tools"

	"github.com/gocraft/dbr"
)

//Inbound used for get Inbound goods
type Inbound struct {
	PoDate            string       `db:"-" json:"po_date" `
	PoDateTime        dbr.NullTime `db:"po_date" json:"-"`
	SKU               string       `db:"sku" json:"sku" `
	PoLine            int          `db:"po_line_id" json:"po_line_id"`
	QtyPO             int          `db:"qty_po" json:"jumlah_pesanan" `
	QtyReceive        int          `db:"qty_receive" json:"jumlah_diterima" `
	PurchasePrice     float64      `db:"harga_beli" json:"harga_beli" `
	TotalInboundPrice float64      `db:"-" json:"total" `
	KwitansiNum       string       `db:"kwitansi_num" json:"nomor_kwitansi" `
	Noted             string       `db:"-" json:"noted" `
	ItemName          string       `db:"item_name" json:"item_name" `
	Rcv               []rcv        `db:"-" json:"rcv" `
}

type rcv struct {
	RcvDate     dbr.NullString `db:"-" json:"receive_date"`
	RcvDateTime dbr.NullTime   `db:"receive_date" json:"-"`
	QtyReceive  dbr.NullInt64  `db:"qty_receive" json:"jumlah_diterima"`
	RcvStatus   dbr.NullString `db:"rcv_status" json:"rcv_status"`
	DoStatus    dbr.NullString `db:"do_status" json:"do_status"`
}

//SearchInbounds used for search data from inbounds goods
func (obj *Inbound) SearchInbounds(db *dbr.Session) (output object.SearchResult) {
	var data []Inbound
	var query string

	tx, err := db.Begin()
	defer tx.RollbackUnlessCommitted()

	if err != nil {
		return
	}

	total := fmt.Sprintf(`
	SELECT COUNT(*) as total
	FROM   po_headers poh 
       JOIN po_lines pol 
         ON poh.po_header_id = pol.po_header_id 
       JOIN inventory inv 
         ON pol.sku = inv.sku 
       JOIN delivery_order do 
         ON do.po_line_id = pol.po_line_id 
       LEFT JOIN receiving rcv 
              ON rcv.do_id = do.do_id 
GROUP  BY poh.po_date, 
          pol.sku, 
          inv.item_name, 
          pol.qty_po, 
          pol.po_line_id, 
          do.po_line_id, 
          pol.purchase_price, 
          do.kwitansi_num `)

	query = fmt.Sprintf(`SELECT poh.po_date, 
	pol.sku, 
	inv.item_name, 
	pol.qty_po, 
	pol.po_line_id, 
	do.po_line_id, 
	Sum(rcv.qty_receive) qty_receive, 
	pol.purchase_price   AS harga_beli, 
	do.kwitansi_num 
FROM   po_headers poh 
	JOIN po_lines pol 
	  ON poh.po_header_id = pol.po_header_id 
	JOIN inventory inv 
	  ON pol.sku = inv.sku 
	JOIN delivery_order do 
	  ON do.po_line_id = pol.po_line_id 
	LEFT JOIN receiving rcv 
		   ON rcv.do_id = do.do_id 
GROUP  BY poh.po_date, 
	   pol.sku, 
	   inv.item_name, 
	   pol.qty_po, 
	   pol.po_line_id, 
	   do.po_line_id, 
	   pol.purchase_price, 
	   do.kwitansi_num `)

	err = tx.SelectBySql(total).LoadOne(&output)
	if output.Total != 0 {
		_, err = tx.SelectBySql(query).Load(&data)
		if err == nil {
			for index := range data {
				rcvQuery := fmt.Sprintf(`SELECT rcv.receive_date,rcv.qty_receive,rcv.status as rcv_status,do.status do_status
				FROM   delivery_order do 
				left  JOIN receiving rcv 
				  ON rcv.do_id = do.do_id 
				  where do.po_line_id = %v
				  order by receive_date asc`, data[index].PoLine)
				_, err = tx.SelectBySql(rcvQuery).Load(&data[index].Rcv)

				data[index].UnmarshalTime()
				data[index].TotalInboundPrice = data[index].PurchasePrice * float64(data[index].QtyPO)
				// var noted string
				for i := range data[index].Rcv {
					if data[index].Rcv[i].RcvStatus.String == "completed" {
						str1 := strconv.Itoa(int(data[index].Rcv[i].QtyReceive.Int64))
						data[index].Noted = data[index].Noted + data[index].Rcv[i].RcvDate.String + " terima " + str1 + "- "
					}
					if data[index].Rcv[i].DoStatus.String != "completed" {
						data[index].Noted = data[index].Noted + "Masih Menunggu " + "- "
					}
				}

			}
			tx.Commit()
			output.Data = data
		}
	}

	return
}

//UnmarshalTime used for convert date to string
func (obj *Inbound) UnmarshalTime() {
	obj.PoDate = tools.TimeToString(obj.PoDateTime.Time, "timestamp")
	for i := range obj.Rcv {
		obj.Rcv[i].RcvDate.String = tools.TimeToString(obj.Rcv[i].RcvDateTime.Time, "timestamp")
	}
}

//MarshalTime used for convert string to date
func (obj *Inbound) MarshalTime() (err error) {
	if obj.PoDate != "" {
		obj.PoDateTime.Time, err = time.Parse("2006-01-02 15:04:05", obj.PoDate)
		if err != nil {
			return
		}
		obj.PoDateTime.Valid = true
	}
	if len(obj.Rcv) != 0 {
		for i := range obj.Rcv {
			obj.Rcv[i].RcvDateTime.Time, err = time.Parse("2006-01-02 15:04:05", obj.Rcv[i].RcvDate.String)
		}
	}

	return
}

//IsInitial used for checking empty struct
func (obj Inbound) IsInitial() bool {
	if reflect.DeepEqual(obj, Inbound{}) {
		return true
	}
	return false
}

//Clear used for clear value
func (obj *Inbound) Clear() {
	p := reflect.ValueOf(obj).Elem()
	p.Set(reflect.Zero(p.Type()))
}
