package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func initDB() (err error) {
	dsn := "root:admin@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True"
	//dsn:="root:admin/test"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	return nil
}

func Insert () {
	s := "insert into user_t (username,passwd) values(?,?)"
	r, err := db.Exec(s, "zhangsan", "zs123")
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}else{
		id, err := r.LastInsertId()
		if err != nil {
			fmt.Printf("last err:%s\n", err)
		}else {
			fmt.Printf("theId:%s\n", id)
		}
	}
}

func Select()  {
	sql := "select * from user_t  where id = ? "
	rows,err := db.Query(sql,1)
	defer rows.Close()
	if err!=nil{
		log.Fatal(err)
	}

	columns, _ := rows.Columns()
	length := len(columns)
	for rows.Next() {
		value := make([][]byte, length)
		columnPointers := make([]interface{}, length)
		for i:=0;i<length;i++ {
			columnPointers[i] = &value[i]
		}
		rows.Scan(columnPointers...)

		data := make(map[string]string)

		for k, v := range value {
			key := columns[k]
			data[key] = string(v)
		}
		/*for i:=0;i<length;i++ {
			columnName := columns[i]
			columnValue := columnPointers[i].(*interface{})
			data[columnName] = string(*columnValue)
		}*/
		log.Print(data)
	}

	//return result
}

func main() {
	/*db, err := sql.Open("mysql", "root:admin@/test")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	fmt.Printf("db:%v\n", db)
	*/
	err := initDB()
	if err != nil {
		fmt.Printf("err！%v\n",err)
	}else{
		fmt.Printf("连接成功！")
	}

	//Insert()
	Select()
	//fmt.Print(myrows)
}