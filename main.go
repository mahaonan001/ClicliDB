package main

import (
	"fmt"
	"reflect"
)

type colum_append interface {
	Append_content(any)
}

func (c *colum) Append_content(contex any) {
	if reflect.TypeOf(contex) != c.Type {
		panic("different type with colum")
	}
	c.content = contex
}

type colum struct {
	name    string
	Type    reflect.Type
	content any
}
type table struct {
	name   string
	colums []colum
}

func (db *DB) Create_table(name string, colums_name []string, colums_type []reflect.Type) {
	len_n := len(colums_name)
	len_t := len(colums_type)
	if len_n != len_t {
		panic("length of colums_name must be the same with colums_type")
	}
	var table table
	table.name = name
	for i := 0; i < len_n; i++ {
		var colum colum
		colum.name = colums_name[i]
		colum.Type = colums_type[i]
		table.colums = append(table.colums, colum)
	}
	db.tables = append(db.tables, table)
}

type DB struct {
	name   string
	tables []table
}

func (dbs *DBS) Create_db(name string) {
	var db = DB{}
	db.name = name
	dbs.databases = append(dbs.databases, db)
}

type DBS struct {
	databases []DB
}
type GonySQL interface {
	Create_db(string)
}
type DataBase interface {
	Create_table(string, []string, []reflect.Type)
}

var dbs DBS

func main() {
	var chiose int
	fmt.Println("欢迎来到GonySQL数据库系统")
	switch chiose {
	case 1:
		var str string
		fmt.Print("请输入数据库名称:")
		fmt.Scan(&str)
		dbs.Create_db(str)
	}
}
