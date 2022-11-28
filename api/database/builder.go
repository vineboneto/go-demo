package database

import (
	"fmt"
	"strings"
)

type Query struct {
	q         string
	listWhere []string
	args      []any
	limit     int
	offset    int
}

func Build() *Query {
	return &Query{}

}

func (q *Query) Where() *Query {
	q.listWhere = append(q.listWhere, " WHERE 1 = 1 ")
	return q
}

func (q *Query) And(s string, v any) *Query {

	if v != "" && v != 0 && v != nil {
		q.listWhere = append(q.listWhere, s)
		q.args = append(q.args, v)
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

func (q *Query) String() (string, []any) {
	strLimit := ""
	strOffset := ""

	if q.limit != 0 {
		strLimit = fmt.Sprintf(" LIMIT %d ", q.limit)
	}

	if q.offset != 0 {
		strOffset = fmt.Sprintf(" OFFSET %d ", q.offset+q.limit-1)
	}

	return strings.Join(q.listWhere, " AND ") + strOffset + strLimit, q.args
}
