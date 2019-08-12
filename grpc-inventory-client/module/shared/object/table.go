package object

import (
	"errors"
	"strings"

	"github.com/technical-assessment/iqbal/salestock/grpc-inventory-client/module/shared/tools"
)

//TableOperation used for parameter for table operation
type TableOperation struct {
	SortBy     string `form:"sort_by"`
	SortMethod string `form:"sort_method"`
	Limit      int    `form:"limit"`
	Offset     int    `form:"offset"`
}

//Valid used for validate 'TableOperation' object
func (obj *TableOperation) Valid(target interface{}) error {
	var exist bool
	if obj.SortBy != "" {
		for _, row := range tools.GetTag("db", target) {
			if row == strings.ToLower(obj.SortBy) {
				exist = true
			}
		}

		if !exist {
			return errors.New("Raise custom validation")
		}
	}

	if obj.SortMethod != "" {
		if strings.ToLower(obj.SortMethod) != "asc" && strings.ToLower(obj.SortMethod) != "desc" {
			return errors.New("Raise custom validation")
		}
	} else {
		obj.SortMethod = "asc"
	}

	if obj.Limit == 0 {
		obj.Limit = 10
	}

	return nil
}
