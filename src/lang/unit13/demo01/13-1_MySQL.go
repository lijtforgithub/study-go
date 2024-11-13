package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 数据库连接信息
	dataSource := "root:admin@tcp(127.0.0.1:3306)/test"

	// 打开数据库连接
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 测试连接
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("数据库连接成功")

	// 执行 SQL 查询
	rows, err := db.Query("SELECT * FROM log")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// 处理查询结果
	for rows.Next() {
		var id int
		var content string
		if err := rows.Scan(&id, &content); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Content: %s\n", id, content)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	// 开始事务
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// 插入数据
	_, err = tx.Exec("INSERT INTO log (content) VALUES (?)", "测试插入")
	if err != nil {
		_ = tx.Rollback()
		log.Fatal(err)
	}

	// 更新数据
	_, err = tx.Exec("UPDATE log SET content = ? WHERE id = ?", "测试更新", 10)
	if err != nil {
		_ = tx.Rollback()
		log.Fatal(err)
	}

	// 提交事务
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Transaction completed successfully!")
}
