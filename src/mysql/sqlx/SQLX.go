package main

// go get github.com/jmoiron/sqlx
import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"math/rand"
	"strconv"
)

type user struct {
	Id       int
	Name     string
	NickName string `db:"nick_name"`
}

const TAB_NAME = "go_user"

func main() {
	db := createConnect()
	defer db.Close()

	//id := insertOne(db)
	//if id == 0 {
	//	id = 1
	//}
	//
	//selectOne(db, id)
	//updateOne(db, id)
	//selectMulti(db)
	//deleteOne(db, id)
	//
	//fmt.Println()
	//id = transaction(db)
	//selectOne(db, id)
	//prepareSql(db)

	nameQuery(db)
	nameExec(db)
}

func nameQuery(db *sqlx.DB) {
	var u = user{
		Id: 1,
	}
	sqlStr := fmt.Sprintf("SELECT * FROM %s WHERE id = :id", TAB_NAME)
	rs, err := db.NamedQuery(sqlStr, &u)

	if err != nil {
		fmt.Println(err)
	} else {
		for rs.Next() {
			_ = rs.StructScan(&u)
			log.Println("nameQuery", u.Id, u.Name, u.NickName)
		}
	}
}

func nameExec(db *sqlx.DB) {
	sqlStr := fmt.Sprintf("UPDATE %s SET name = :name, nick_name = :nick_name WHERE id = :id", TAB_NAME)
	i := strconv.Itoa(rand.Intn(100))

	_, err := db.NamedExec(sqlStr, user{
		Id:       1,
		Name:     "namedExec" + i,
		NickName: "namedExec" + i,
	})
	if err != nil {
		fmt.Println(err)
	}

	_, err = db.NamedExec(sqlStr, map[string]any{
		"id":        2,
		"name":      "namedExec" + i,
		"nick_name": "namedExec" + i,
	})
	if err != nil {
		fmt.Println(err)
	}
}

func prepareSql(db *sqlx.DB) {
	sqlStr := fmt.Sprintf("UPDATE %s SET name = ?, nick_name = ? WHERE id = ?", TAB_NAME)
	st, err := db.Preparex(sqlStr)
	if err != nil {
		panic(err)
	}

	i := strconv.Itoa(rand.Intn(100))
	rs1 := st.MustExec("p-name"+i, "p-nickName"+i, 1)
	rs2 := st.MustExec("p-name"+i, "p-nickName"+i, 2)

	n1, _ := rs1.RowsAffected()
	n2, _ := rs2.RowsAffected()
	log.Println("预执行", n1, n2)
}

func transaction(db *sqlx.DB) int64 {
	tx := db.MustBegin()

	sqlStr := fmt.Sprintf("INSERT INTO %s (name, nick_name) VALUES(?, ?)", TAB_NAME)
	rs := tx.MustExec(sqlStr, "tx-name", "tx-nickName")

	id, _ := rs.LastInsertId()
	log.Println("事务插入数据", id)

	err := tx.Rollback()
	if err != nil {
		fmt.Println("事务回滚失败", err)
	}

	return id
}

func insertOne(db *sqlx.DB) int64 {
	sqlStr := fmt.Sprintf("INSERT INTO %s (name, nick_name) VALUES(?, ?)", TAB_NAME)
	rs := db.MustExec(sqlStr, "i-name", "i-nickName")

	id, _ := rs.LastInsertId()
	log.Println("插入数据", id)
	return id
}

func updateOne(db *sqlx.DB, id int64) {
	sqlStr := fmt.Sprintf("UPDATE %s SET name = ?, nick_name = ? WHERE id = ?", TAB_NAME)
	rs := db.MustExec(sqlStr, "u-name", "u-nickName", id)

	n, _ := rs.RowsAffected()
	log.Println("更新数据", id, n)
}

func deleteOne(db *sqlx.DB, id int64) {
	sqlStr := fmt.Sprintf("DELETE FROM %s WHERE id = ?", TAB_NAME)
	rs := db.MustExec(sqlStr, id)

	n, _ := rs.RowsAffected()
	log.Println("删除数据", id, n)
}

func selectOne(db *sqlx.DB, id int64) {
	var u user
	sqlStr := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", TAB_NAME)
	err := db.Get(&u, sqlStr, id)
	if err != nil {
		fmt.Println(err)
	} else {
		log.Println("查询单个", u.Id, u.Name, u.NickName)
	}
}

func selectMulti(db *sqlx.DB) {
	var us []user
	sqlStr := fmt.Sprintf("SELECT id, nick_name, name FROM %s", TAB_NAME)
	err := db.Select(&us, sqlStr)

	if err != nil {
		fmt.Println(err)
	} else {
		for _, v := range us {
			log.Println("查询多个", v.Id, v.Name, v.NickName)
		}
	}
}

func createConnect() *sqlx.DB {
	dataSource := "root:admin@tcp(127.0.0.1:3306)/test"
	db := sqlx.MustConnect("mysql", dataSource)

	err := db.Ping()
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(1)
	fmt.Println("数据库连接成功")
	return db
}
