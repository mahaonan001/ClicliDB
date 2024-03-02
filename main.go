package main

import (
	"fmt"
	"reflect"
)

// type colum_append interface {
// 	Append_content(any)
// }

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
func (db *DB) In_dbs() {
	var chiose int
	for {
		fmt.Println(`
		0:返回上一级
		1:创建表
		2:加载表
		`)
		fmt.Scan(&chiose)
		switch chiose {
		case 0:
			return
		case 1:
			var str string
			var columnsName []string
			var columnsType []reflect.Type

			fmt.Print("请输入表名: ")
			fmt.Scan(&str)

			// 读取字符串切片
			fmt.Print("Enter the number of columns: ")
			var numColumns int
			fmt.Scan(&numColumns)

			// 初始化切片
			columnsName = make([]string, numColumns)
			columnsType = make([]reflect.Type, numColumns)

			// 读取每个列的名称
			for i := 0; i < numColumns; i++ {
				fmt.Printf("Enter column name %d: ", i)
				fmt.Scan(&columnsName[i])
			}

			// 注意：fmt.Scan不支持直接读取到reflect.Type切片
			// 你需要根据实际情况来处理columnsType的赋值
			// 这里我们假设columnsType是一个固定类型的切片，例如都是string类型
			for i := 0; i < numColumns; i++ {
				columnsType[i] = reflect.TypeOf("")
			}

			// 输出读取的列名
			fmt.Println("Column names:")
			for _, name := range columnsName {
				fmt.Println(name)
			}

			// 输出读取的列类型
			fmt.Println("Column types:")
			for _, t := range columnsType {
				fmt.Println(t)
			}
			db.Create_table(str, columnsName, columnsType)
		default:
			fmt.Println("无效操作")
		}
	}

}

type DB struct {
	name   string
	tables []table
}

func (dbs *DBS) Create_db(name string) DB {
	var db DB
	db.name = name
	dbs.databases = append(dbs.databases, db)
	return db
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

	fmt.Println("		欢迎来到GonySQL数据库系统")

	for {
		fmt.Println(`
		请输入数字:
			0:退出GonySQL
			1:创建数据库
			2:查看SQL
			3:载入数据库`)
		fmt.Scan(&chiose)
		switch chiose {
		case 1:
			var str string
			fmt.Print("请输入数据库名称:")
			fmt.Scan(&str)
			db := dbs.Create_db(str)
			db.In_dbs()
		case 2:
			fmt.Println(dbs)
		case 3:
			fmt.Println()
			fmt.Println("请选择数据库:")
		case 0:
			return
		default:
			fmt.Println("无效操作")
		}
	}
}
