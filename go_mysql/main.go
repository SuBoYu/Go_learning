package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	USERNAME = "root"
	PASSWORD = "tonysu518"
	NETWORK  = "tcp"
	SERVER   = "127.0.0.1"
	PORT     = 3306
	DATABASE = "demo"
)

type User struct {
	ID       string
	Username string
	Password string
}

func CreateTable(db *sql.DB) error {
	sql := `CREATE TABLE IF NOT EXISTS users(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
			username VARCHAR(64),
			password VARCHAR(64)
		);`
	if _, err := db.Exec(sql); err != nil {
		fmt.Println("建立table發生錯誤: ", err)
		return err
	}
	fmt.Println("建立Table成功!")
	return nil
}

func InsertUser(DB *sql.DB, username, password string) error {
	_, err := DB.Exec("insert INTO users(username,password) values(?,?)", username, password)
	if err != nil {
		fmt.Printf("建立使用者失敗，原因是：%v", err)
		return err
	}
	fmt.Println("建立使用者成功！")
	return nil
}

func QueryUser(db *sql.DB, username string) {
	user := new(User)
	row := db.QueryRow("Select * from users where username=? ", username)
	if err := row.Scan(&user.ID, &user.Username, &user.Password); err != nil {
		fmt.Println("映射使用者失敗，原因為：%v\n", err)
		return
	}
	fmt.Println("查詢使用者成功", *user)
}

func main() {

	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)

	db, err := sql.Open("mysql", conn)
	if err != nil {
		fmt.Println("開啟 MYSQL 連線發生錯誤，原因為: ", err)
		return
	}
	if err := db.Ping(); err != nil {
		fmt.Println("資料庫連線錯誤，原因為: ", err.Error())
		return
	}
	fmt.Println("連線成功")
	defer db.Close()

	QueryUser(db, "test")
}
