package utils

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

	modalDb.db = db
	modalDb.m = &Modal{}
	modalDb.m.tabelName = url
	if (!modalDb.db.HasTable(url)) {
		modalDb.CreateTable()
	}

	return modalDb
}

func (u Modal) TableName() string {
    return u.tabelName
}

func (mb * ModalDb) CreateTable() {
	err := mb.db.CreateTable(mb.m).Error
	if nil != err {
		fmt.Println(err)
	}
}

// 插入数据
func (mb * ModalDb) InsertData(url string, email string) {
	mb.m.Url = url
	mb.m.Email = email
	err := mb.db.Create(mb.m).Error
	if nil != err {
		fmt.Println(err)
	}
}

func (mb * ModalDb) SelectData(num int) ([]Modal, error) {
	var emailModals []Modal
	err := mb.db.Limit(num).Table(mb.m.tabelName).
		Where("status=false").
		Where("email like ? or email like ? or email like ? or email like ? or email like ?", 
		"%126%", "%163%", "%qq%", "%hotmail%", "%sina%").
		Find(&emailModals).Error
	return emailModals, err
}

func (mb * ModalDb) UpdateStatus(email string, status bool) error {
	err := mb.db.Model(&Modal{}).Table(mb.m.tabelName).
		Where("email = ?", email).
		Update("status", status).Error
	return err
}

func (mb * ModalDb) Close() {
	mb.db.Close()
}