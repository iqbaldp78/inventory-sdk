package crud

import (
	"fmt"
	"reflect"

	"github.com/technical-assessment/iqbal/salestock/grpc-inventory/module/shared/object"

	"github.com/gocraft/dbr"
)

//OnhandQTY used for get onhand qty in inventory
type OnhandQTY struct {
	SKU      string `db:"sku" json:"sku" binding:"required"`
	ItemName string `db:"item_name" json:"item_name" binding:"required"`
	Stock    int    `db:"stock" json:"stock"`
}

//Search used for search data from table OnhandQTY
func (obj *OnhandQTY) Search(db *dbr.Session) (output object.SearchResult) {
	var data []OnhandQTY

	total := fmt.Sprintf("SELECT COUNT(*) as total FROM inventory")
	query := fmt.Sprintf("SELECT * FROM inventory ORDER BY created_on desc")

	tx, err := db.Begin()
	defer tx.RollbackUnlessCommitted()

	if err != nil {
		return
	}

	err = tx.SelectBySql(total).LoadOne(&output)

	if output.Total != 0 {
		_, err = tx.SelectBySql(query).Load(&data)
		if err == nil {
			tx.Commit()
			output.Data = data
		}
	}

	return

}

//IsInitial used for checking empty struct
func (obj OnhandQTY) IsInitial() bool {
	if reflect.DeepEqual(obj, OnhandQTY{}) {
		return true
	}
	return false
}

//Clear used for clear value
func (obj *OnhandQTY) Clear() {
	p := reflect.ValueOf(obj).Elem()
	p.Set(reflect.Zero(p.Type()))
}
