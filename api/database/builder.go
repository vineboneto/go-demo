package database

import (
	"strings"
)

type Query struct {
	raw         string
	prefixWhere []string
	args        []any
	limit       int
	offset      int
}

func Build() *Query {
	return &Query{}

}

func (q *Query) Raw(raw string) *Query {
	q.raw = raw
	return q
}

func (q *Query) Where() *Query {
	q.prefixWhere = append(q.prefixWhere, " WHERE 1 = 1 ")
	return q
}

func (q *Query) And(s string, v any) *Query {

	if v != "" && v != 0 && v != nil {
		q.prefixWhere = append(q.prefixWhere, s)
		q.args = append(q.args, v)
	}
	return q
}

func (q *Query) AndLike(s string, v string) *Query {

	if v != "" {
		q.prefixWhere = append(q.prefixWhere, s)
		q.args = append(q.args, "%"+v+"%")
	}
	return q
}

func (q *Query) Offset(v int) *Query {
	if v > 0 {
		q.offset = v
	}
	return q
}

func (q *Query) Limit(v int) *Query {
	if v > 0 {
		q.limit = v
	}
	return q
}

func (q *Query) String() (string, []any) {
	rawLimit := ""
	rawOffset := ""

	if q.limit > 0 {
		q.args = append(q.args, q.limit)
		rawLimit = " LIMIT ? "
	}
	if q.offset > 0 {
		q.args = append(q.args, q.offset+q.limit-1)
		rawOffset = " OFFSET ? "
	}

	return q.raw + strings.Join(q.prefixWhere, " AND ") + rawLimit + rawOffset, q.args
}
