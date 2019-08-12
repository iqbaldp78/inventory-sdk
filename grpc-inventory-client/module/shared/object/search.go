package object

import "reflect"

//SearchParam used as parameter for search data
type SearchParam struct {
	Key    string `form:"key"`
	Limit  int    `form:"limit"`
	Offset int    `form:"offset"`
}

//SearchParamDate used as parameter for search by date data
type SearchParamDate struct {
	DateFrom string `form:"date_from" binding:"required"`
	DateTo   string `form:"date_to" binding:"required"`
}

//SearchResult used as result for search data
type SearchResult struct {
	Total int         `db:"total" json:"total,omitempty"`
	Data  interface{} `db:"-" json:"data,omitempty"`
}

//IsInitial used for check this object empty or not
func (obj SearchResult) IsInitial() bool {
	if reflect.DeepEqual(obj, SearchResult{}) {
		return true
	}
	return false
}
