package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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

func insert () {
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

	insert()
}