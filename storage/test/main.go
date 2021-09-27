package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "Alienegra"
	password = "yyh851521"
	dbname   = "dbname"
)

type Teacher struct {
	ID  int
	Age int
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

	//建表
	_, err = db.Exec(`
CREATE TABLE IF NOT EXISTS user_presence_data (
	user_id INT NOT NULL,
	device INT NOT NULL,
	token TEXT NOT NULL,
	created_time BIGINT NOT NULL
);
CREATE UNIQUE INDEX IF NOT EXISTS user_id_index ON user_presence_data(user_id, device);
`)

	_, err = db.Exec(`
CREATE TABLE IF NOT EXISTS user_data (
	user_id serial PRIMARY KEY,
    user_name TEXT NOT NULL,
    password TEXT NOT NULL,
    phone_num TEXT NOT NULL,
    email TEXT NOT NULL,
    gender INT NOT NULL,
    age INT NOT NULL,
    phone_ver_flag INT,
    mail_ver_flag INT,
    google_ver_flag INT,
    google_key TEXT,
	gesture_ver_flag INT,
	gesture_key TEXT,
	finger_ver_flag INT,
	finger_key TEXT,
	fish_ver_flag INT,
	fish_code TEXT,
	certify_num TEXT,
	certify_front TEXT,
	certify_back TEXT,
	user_real_name TEXT,
    created_ts BIGINT NOT NULL
);
`)
	if err != nil {
		panic(err)
	}
	//插入数据
	//stmt, err := db.Prepare("INSERT INTO userinfo(username,departname,created) VALUES($1,$2,$3) RETURNING uid")
	//if err != nil {
	//	panic(err)
	//}
	//
	//_, err = stmt.Exec("冒险王", "游戏部门", "2012-12-09")
	//if err != nil {
	//	panic(err)
	//}

	//查询数据
	//rows, _ := db.Query("SELECT * FROM userinfo")
	//
	//for rows.Next() {
	//	var uid int
	//	var username string
	//	var department string
	//	var created string
	//	err = rows.Scan(&uid, &username, &department, &created)
	//	fmt.Println(uid)
	//	fmt.Println(username)
	//	fmt.Println(department)
	//	fmt.Println(created)
	//}

	//pg不支持这个函数，因为他没有类似MySQL的自增ID
	//id, err := res.LastInsertId()
	//if err != nil {
	//	panic(err)
	//}

	//fmt.Println(id)
}
