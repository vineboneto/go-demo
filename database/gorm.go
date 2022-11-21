package main

import (
	"fmt"
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Person struct {
	FirstName string
	LastName  string
	Email     string
}

type Cliente struct {
	Codcli int
	Nomcli string
}

type Query struct {
	Q         string
	ListWhere []string
}

func Build() *Query {
	return &Query{}

}

func (q *Query) Select(s string) *Query {
	q.Q = " SELECT " + s + " "
	return q
}

func (q *Query) From(s string) *Query {
	q.Q = q.Q + " FROM " + s
	return q
}

func (q *Query) Where() *Query {
	q.ListWhere = append(q.ListWhere, " WHERE 1 = 1 ")
	return q
}

func (q *Query) And(s string, v any) *Query {

	if v != "" && v != nil {
		q.ListWhere = append(q.ListWhere, fmt.Sprintf(s, v))
	}
	return q
}

func (q *Query) ToString() string {
	return q.Q + strings.Join(q.ListWhere, " AND ")
}

func timeTrack(start time.Time) {
	fmt.Println("total: ", time.Since(start))
}

func main() {

	defer timeTrack(time.Now())

	db, _ := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=localhost user=postgres password=1234 dbname=godemo port=5432 sslmode=disable",
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	var persons []Person

	db.Table("person").Select("*").Where(&Person{FirstName: "Jason"}).Limit(500).Find(&persons)

	firstName := "Jason"
	var lastName string

	sql := Build().Where().And("first_name = '%s'", firstName).And("last_name = '%s'", lastName).ToString()

	fmt.Println(sql)

	// db.Raw(sql).Scan(&persons)

	fmt.Println(persons)

	// db2, _ := gorm.Open(sqlserver.Open("url"), &gorm.Config{})

	// var clientes []Cliente

	// db2.Select("codcli, nomcli").Table("e085cli").Find(&clientes)

	// fmt.Println(len(clientes))
}
