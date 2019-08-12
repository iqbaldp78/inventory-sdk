package object

import "reflect"

//SearchParam used as parameter for search data
type SearchParam struct {
	Key    string `form:"key"`
	Limit  int64  `form:"limit" binding:"required,lte=100"`
	Offset int64  `form:"offset"`
}

//SearchResult used as result for search data
type SearchResult struct {
	Total int64       `db:"total" json:"total,omitempty"`
	Data  interface{} `db:"-" json:"data,omitempty"`
}

//IsInitial used for check this object empty or not
func (obj SearchResult) IsInitial() bool {
	if reflect.DeepEqual(obj, SearchResult{}) {
		return true
	}
	return false
}
