package markdown

import (
	"github-stars/schemas"
	"reflect"
	"sort"
)

type By func(p1, p2 *schemas.GitHubResponseField) bool

func Prop(field string, asc bool) func(p1, p2 *schemas.GitHubResponseField) bool {
	return func(p1, p2 *schemas.GitHubResponseField) bool {

		v1 := reflect.Indirect(reflect.ValueOf(p1)).FieldByName(field)
		v2 := reflect.Indirect(reflect.ValueOf(p2)).FieldByName(field)

		ret := false

		switch v1.Kind() {
		case reflect.Int64:
			ret = v1.Int() < v2.Int()
		case reflect.Float64:
			ret = v1.Float() < v2.Float()
		case reflect.String:
			ret = v1.String() < v2.String()
		}

		if asc {
			return ret
		}
		return !ret
	}
}

func (by By) Sort(responseFields []schemas.GitHubResponseField) {
	ps := &responseFieldsSorter{
		responseFields: responseFields,
		by:      by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(ps)
}

// Len is part of sort.Interface.
func (s *responseFieldsSorter) Len() int {
	return len(s.responseFields)
}

// Swap is part of sort.Interface.
func (s *responseFieldsSorter) Swap(i, j int) {
	s.responseFields[i], s.responseFields[j] = s.responseFields[j], s.responseFields[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *responseFieldsSorter) Less(i, j int) bool {
	return s.by(&s.responseFields[i], &s.responseFields[j])
}


type responseFieldsSorter struct {
	responseFields []schemas.GitHubResponseField
	by      func(p1, p2 *schemas.GitHubResponseField) bool // Closure used in the Less method.
}