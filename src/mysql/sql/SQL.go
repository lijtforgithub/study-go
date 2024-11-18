package main

// github.com/go-sql-driver/mysql
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"math/rand"
	"strconv"
)

type user struct {
	Id       int
	Name     string
	nickName string
}

const TAB_NAME = "go_user"

func main() {
	db := createConnect()
	defer db.Close()

	// 保存一个
	id := insertOne(db)
	if id == 0 {
		id = 1
	}

	// 查询单个
	selectOne(db, id)
	// 更新一个
	updateOne(db, id)
	// 查询多个
	selectMulti(db)
	// 删除一个
	deleteOne(db, id)

	// 事务
	fmt.Println()
	id = transaction(db)
	selectOne(db, id)
	// 预执行
	prepareSql(db)
}

func prepareSql(db *sql.DB) {
	sqlStr := fmt.Sprintf("UPDATE %s SET name = ?, nick_name = ? WHERE id = ?", TAB_NAME)
	st, err := db.Prepare(sqlStr)
	if err != nil {
		panic(err)
	}

	i := strconv.Itoa(rand.Intn(100))
	rs1, _ := st.Exec("p-name"+i, "p-nickName"+i, 1)
	rs2, _ := st.Exec("p-name"+i, "p-nickName"+i, 2)

	n1, _ := rs1.RowsAffected()
	n2, _ := rs2.RowsAffected()
	log.Println("预执行", n1, n2)
}

func transaction(db *sql.DB) int64 {
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	sqlStr := fmt.Sprintf("INSERT INTO %s (name, nick_name) VALUES(?, ?)", TAB_NAME)
	rs, err := tx.Exec(sqlStr, "tx-name", "tx-nickName")

	if err != nil {
		fmt.Println(err)
	} else {
		id, _ := rs.LastInsertId()
		log.Println("事务插入数据", id)

		err = tx.Rollback()
		if err != nil {
			fmt.Println("事务回滚失败", err)
		}

		return id
	}

	return 0
}

func insertOne(db *sql.DB) int64 {
	sqlStr := fmt.Sprintf("INSERT INTO %s (name, nick_name) VALUES(?, ?)", TAB_NAME)
	rs, err := db.Exec(sqlStr, "i-name", "i-nickName")
	if err != nil {
		fmt.Println(err)
	} else {
		id, _ := rs.LastInsertId()
		n, _ := rs.RowsAffected()
		log.Println("插入数据", id, n)
		return id
	}

	return 0
}

func updateOne(db *sql.DB, id int64) {
	sqlStr := fmt.Sprintf("UPDATE %s SET name = ?, nick_name = ? WHERE id = ?", TAB_NAME)
	rs, err := db.Exec(sqlStr, "u-name", "u-nickName", id)
	if err != nil {
		fmt.Println(err)
	} else {
		_, _ = rs.LastInsertId()
		n, _ := rs.RowsAffected()
		log.Println("更新数据", id, n)
	}
}

func deleteOne(db *sql.DB, id int64) {
	sqlStr := fmt.Sprintf("DELETE FROM %s WHERE id = ?", TAB_NAME)
	rs, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Println(err)
	} else {
		_, _ = rs.LastInsertId()
		n, _ := rs.RowsAffected()
		log.Println("删除数据", id, n)
	}
}

func selectOne(db *sql.DB, id int64) {
	var u user
	sqlStr := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", TAB_NAME)
	err := db.QueryRow(sqlStr, id).Scan(&u.Id, &u.Name, &u.nickName)
	if err != nil {
		fmt.Println(err)
	} else {
		log.Println("查询单个", u.Id, u.Name, u.nickName)
	}
}

func selectMulti(db *sql.DB) {
	sqlStr := fmt.Sprintf("SELECT id, nick_name, name FROM %s", TAB_NAME)
	rows, err := db.Query(sqlStr)

	if err != nil {
		fmt.Println(err)
	} else {
		defer rows.Close()
		var id int64
		var nickName, name string
		for rows.Next() {
			_ = rows.Scan(&id, &nickName, &name)
			log.Println("查询多个", id, name, nickName)
		}
	}
}

func createConnect() *sql.DB {
	dataSource := "root:admin@tcp(127.0.0.1:3306)/test"
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(1)
	fmt.Println("数据库连接成功")
	return db
}
