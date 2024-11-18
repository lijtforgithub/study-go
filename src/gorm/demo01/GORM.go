package main

// go get -u gorm.io/gorm
// go get -u gorm.io/driver/mysql
import (
	"log"
	orm "study-go/src/gorm"
)

func main() {
	db := orm.CreateConnect()

	u := orm.GoUser{
		Id:       1,
		Name:     "GORM",
		NickName: "gorm",
	}

	_ = db.Create(&u)
	log.Println(u.Id)

	var u1 orm.GoUser
	db.First(&u1, u.Id)
	log.Println(u1.Name)

	db.Model(&u1).Update("Name", "u-name")
	db.Model(&u1).Updates(orm.GoUser{NickName: "u-nickName"})

	db.Delete(&u1, u1.Id)
}
