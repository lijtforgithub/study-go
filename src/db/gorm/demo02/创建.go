package main

import (
	"fmt"
	"gorm.io/gorm"
	orm "study-go/src/db/gorm"
)

func main() {
	db := orm.CreateConnect()

	//saveOne(db)
	//saveMulti(db)

	saveBatch(db)

	// 用指定的字段创建记录
	//saveSelect(db)
	//saveOmit(db)

}

func saveOne(db *gorm.DB) {
	u := orm.User{
		Name:     "GORM",
		NickName: "gorm",
	}

	result := db.Create(&u)
	// 返回主键
	fmt.Println(u.Id)
	// 返回插入记录行数
	fmt.Println(result.RowsAffected)
}

func saveMulti(db *gorm.DB) {
	us := []*orm.User{
		{Name: "GORM-1", NickName: "gorm-1"},
		{Name: "GORM-2", NickName: "gorm-2"},
	}

	result := db.Create(us)
	fmt.Println(result.RowsAffected)
}

func saveBatch(db *gorm.DB) {
	us := []orm.User{
		{Name: "GORM-1", NickName: "gorm-1"},
		{Name: "GORM-2", NickName: "gorm-2"},
		{Name: "GORM-3", NickName: "gorm-3"},
	}

	_ = db.CreateInBatches(&us, 2)
}

func saveSelect(db *gorm.DB) {
	u := orm.User{
		Name:     "Select-Name",
		NickName: "Select-NickName",
	}

	_ = db.Select("Name").Create(&u)

	var u1 orm.User
	db.First(&u1, orm.User{Id: u.Id})
	fmt.Println(u1.Id, u1.Name, u1.NickName)
}

func saveOmit(db *gorm.DB) {
	u := orm.User{
		Name:     "Omit-Name",
		NickName: "Omit-NickName",
	}

	_ = db.Omit("Name").Create(&u)

	var u1 orm.User
	db.First(&u1, orm.User{Id: u.Id})
	fmt.Println(u1.Id, u1.Name, u1.NickName)
}
