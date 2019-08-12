package crud

import (
	"fmt"
	"reflect"
	"time"

	"github.com/gocraft/dbr"

	"github.com/technical-assessment/iqbal/salestock/grpc-inventory/module/shared/object"
	"github.com/technical-assessment/iqbal/salestock/grpc-inventory/module/shared/tools"
)

//Outbound used for get outbounds goods
type Outbound struct {
	OrderDate          string        `db:"-" json:"order_date" binding:"required"`
	OrderDateTime      dbr.NullTime  `db:"order_date" json:"-" binding:"required"`
	Sku                string        `db:"sku" json:"sku" binding:"required"`
	Itemname           string        `db:"item_name" json:"item_name" binding:"required"`
	QtyOut             dbr.NullInt64 `db:"ship_qty" json:"ship_qty" binding:"required"`
	SellingPrice       float32       `db:"selling_price" json:"selling_price" binding:"required"`
	TotalOutboundPrice float32       `db:"-" json:"total_selling_price" binding:"required"`
	Noted              string        `db:"-" json:"catatan" binding:"required"`
	OrderNum           string        `db:"order_num" json:"-" binding:"required"`
	ShipStatus         string        `db:"ship_status" json:"-" binding:"required"`
}

//SearchOutbounds used for interact to DB. to get outbound data
func (obj *Outbound) SearchOutbounds(db *dbr.Session) (output object.SearchResult) {
	var data []Outbound

	tx, err := db.Begin()
	defer tx.RollbackUnlessCommitted()
	if err != nil {
		return
	}

	total := fmt.Sprintf(`SELECT COUNT(*) as total FROM sales_order so JOIN so_lines sol on so.order_num = sol.order_num JOIN inventory inv on sol.sku=inv.sku JOIN shipping shp on shp.so_line_id = sol.so_line_id`)
	err = tx.SelectBySql(total).LoadOne(&output)
	query := fmt.Sprintf(`SELECT so.order_date, sol.sku, inv.item_name, sol.selling_price, shp.status ship_status, shp.qty_ship ship_qty,so.order_num FROM sales_order so JOIN so_lines sol ON so.order_num = sol.order_num JOIN inventory inv ON sol.sku = inv.sku JOIN shipping shp ON shp.so_line_id = sol.so_line_id`)

	if output.Total != 0 {
		_, err = tx.SelectBySql(query).Load(&data)
		if err == nil {
			tx.Commit()
			for index := range data {
				data[index].UnmarshalTime()
				if data[index].ShipStatus != "completed" {
					data[index].TotalOutboundPrice = -1
					data[index].SellingPrice = -1
					data[index].Noted = data[index].ShipStatus
				} else {
					data[index].TotalOutboundPrice = data[index].SellingPrice * float32(data[index].QtyOut.Int64)
					data[index].Noted = "Pesanan" + data[index].OrderNum
				}

			}
			output.Data = data
		}
	}
	return output
}

//UnmarshalTime used for convert date to string
func (obj *Outbound) UnmarshalTime() {
	obj.OrderDate = tools.TimeToString(obj.OrderDateTime.Time, "timestamp")
}

//MarshalTime used for convert string to date
func (obj *Outbound) MarshalTime() (err error) {
	if obj.OrderDate != "" {
		obj.OrderDateTime.Time, err = time.Parse("2006-01-02 15:04:05", obj.OrderDate)
		if err != nil {
			return
		}
		obj.OrderDateTime.Valid = true
	}

	return
}

//IsInitial used for checking empty struct
func (obj Outbound) IsInitial() bool {
	if reflect.DeepEqual(obj, Outbound{}) {
		return true
	}
	return false
}

//Clear used for clear value
func (obj *Outbound) Clear() {
	p := reflect.ValueOf(obj).Elem()
	p.Set(reflect.Zero(p.Type()))
}
