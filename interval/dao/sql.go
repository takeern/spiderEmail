package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"fmt"
)

//user表结构体定义
type Modal struct {
	Email string `gorm:"type:varchar(100);not null;unique"` //not null;unique
	Url string `gorm:"type:varchar(100);not null"`
	Status bool   `gorm:"type:varchar(100);default:'false'`      // true 已经发送
	tabelName string `gorm:"-"`
}

type ModalDb struct {
	db *gorm.DB
	m *Modal
}

type SpiderUrl struct {
	Url string `gorm:"type:varchar(100);not null;unique;"`
}

func NewDb(url string) (*ModalDb) {
	modalDb := &ModalDb {}
	db, err := gorm.Open("mysql", "root:maskTakeern@tcp(47.103.12.134:3306)/spider?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	// db.DB().SetMaxIdleConns(10)
	// db.DB().SetMaxOpenConns(100)
	spiderUrl := &SpiderUrl{
		Url: url,
	}
	err = db.Create(&spiderUrl).Error
	fmt.Println(err)

	modalDb.db = db
	modalDb.m = &Modal{}
	modalDb.m.tabelName = url
	createTable(modalDb)

	return modalDb
}

func (u Modal) TableName() string {
    return u.tabelName
}

func createTable(mb *ModalDb) {
	err := mb.db.CreateTable(mb.m).Error
	fmt.Println(err)
}

// 插入数据
func (mb * ModalDb) InsertData(url string, email string) {
	mb.m.Url = url
	mb.m.Email = email
	err := mb.db.Create(mb.m).Error
	fmt.Println(err)
}
