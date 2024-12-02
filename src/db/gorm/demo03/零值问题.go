package main

import (
	"database/sql"
	orm "study-go/src/db/gorm"
)

// 在使用 GORM 进行数据库操作时，有时会遇到零值（如空字符串、数字类型的0等）不被更新的问题。这是因为 GORM 默认只更新非零值字段，即只有当字段的值不是其类型的零值时才会被更新到数据库中。
type zeroValue1 struct {
	Id     int
	Age    int
	Name   string
	Female bool
}

type zeroValue2 struct {
	Id     int
	Age    *int
	Name   sql.NullString
	Female sql.NullBool
}

func main() {
	db := orm.CreateConnect()
	tableName := "go_zero_value"

	zv1 := zeroValue1{
		Age:    0,
		Name:   "",
		Female: false,
	}

	db.Table(tableName).AutoMigrate(&zeroValue1{})
	db.Table(tableName).Create(&zv1)

	//db.Table(tableName).Where("id=1").Update("Age", 0)
	zv1.Name = "u-name"
	// UPDATE `go_zero_value` SET `name`='u-name' WHERE `id` = 9
	db.Table(tableName).Updates(&zv1)

	zero := 0
	zv2 := zeroValue2{
		Id:  zv1.Id,
		Age: &zero,
		Name: sql.NullString{
			String: "",
			Valid:  true,
		},
		Female: sql.NullBool{
			Bool:  false,
			Valid: true,
		},
	}
	// UPDATE `go_zero_value` SET `age`=0,`name`='',`female`=false WHERE `id` = 9
	db.Table(tableName).Updates(&zv2)
}
