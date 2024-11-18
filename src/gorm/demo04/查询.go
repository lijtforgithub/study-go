package main

import (
	"fmt"
	"gorm.io/gorm"
	orm "study-go/src/gorm"
)

func main() {
	db := orm.CreateConnect()

	selectOne(db)
	//selectMulti(db)
}

/*
 * 只有在目标 struct 是指针或者通过 db.Model() 指定 model 时，方法才有效
 */
func selectOne(db *gorm.DB) {
	u1 := orm.GoUser{}
	db.Where("id=1").First(&u1)

	// 根据主键查找
	u2 := orm.GoUser{}
	db.Take(&u2, 2)

	result := map[string]interface{}{}
	u3 := orm.GoUser{
		Id: 3,
	}
	// 设置查询条件 SELECT * FROM `go_user` ORDER BY `go_user`.`id` DESC LIMIT 1
	db.Model(u3).Last(&result)

	var u4 = orm.GoUser{
		Id: 4,
	}
	db.Last(&u4)
}

func selectMulti(db *gorm.DB) {
	var us []orm.GoUser
	db.Find(&us, 1, 2, 3)
	fmt.Println(len(us))
}
