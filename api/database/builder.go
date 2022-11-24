package database

import (
	"fmt"
	"strings"
)

type Query struct {
	q         string
	listWhere []string
	limit     int
	offset    int
}

func Build() *Query {
	return &Query{}

}

func (q *Query) Select(s string) *Query {
	q.q = " SELECT " + s + " "
	return q
}

func (q *Query) From(s string) *Query {
	q.q = q.q + " FROM " + s
	return q
}

func (q *Query) Where() *Query {
	q.listWhere = append(q.listWhere, " WHERE 1 = 1 ")
	return q
}

func (q *Query) And(s string, v any) *Query {

	if v != "" && v != 0 && v != nil {
		q.listWhere = append(q.listWhere, fmt.Sprintf(s, v))
	}
	return q
}

func (q *Query) Offset(v int) *Query {
	if v != 0 {
		q.offset = v
	}
	return q
}

func (q *Query) Limit(v int) *Query {
	if v != 0 {
		q.limit = v
	}
	return q
}

func (q *Query) AndLike(s string, v string) *Query {
	if v != "" {
		q.listWhere = append(q.listWhere, fmt.Sprintf(s, "%"+v+"%"))
	}
	return q
}

func (q *Query) String() string {
	strLimit := ""
	strOffset := ""

	if q.limit != 0 {
		strLimit = fmt.Sprintf(" LIMIT %d ", q.limit)
	}

	if q.offset != 0 {
		strOffset = fmt.Sprintf(" OFFSET %d ", q.offset+q.limit-1)
	}

	return q.q + strings.Join(q.listWhere, " AND ") + strOffset + strLimit
}
