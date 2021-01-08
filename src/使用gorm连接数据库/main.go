package main

import (
	"time"
)

type Area struct {
	Id       int       `gorm:"primary_key,column:id"`
	AreaName string    `gorm:"column:area_name"`
	CityId   int       `gorm:"column:city_id"`
	UserId   int       `gorm:"column:user_id"`
	UpdateAt time.Time `gorm:"column:update_at"`
	CreateAt time.Time `gorm:"column:create_at"`
	DeleteAt time.Time `gorm:"column:delete_at"`
}

//func main__() {
//	//db, _ := gorm.Open("mysql", "root:123456@/test?charset=utf8&parseTime=True&loc=Local")
//	//defer db.Close()
//	//db.LogMode(true)
//
//	//查询一条
//	//db.Table("area").First(area)
//
//	//新增一条
//	//area := Area{AreaName: "礼顿道19号", CityId: 4, UserId: 4,UpdateAt:time.Now(),CreateAt:time.Now()}
//	//rows := db.Table("area").Create(&area)
//
//	//查询所有AreaName为"礼顿道19号"的数据
//	var areas []Area
//	db.Table("area").Where("area_name = ?", "礼顿道19号").Find(&areas)
//	fmt.Println(areas)
//
//	//使用主键获取记录
//	var area Area
//	db.Table("area").First(&area, 4)
//	fmt.Println(area)
//
//	//更新一条数据
//	area.AreaName = "礼顿道20号"
//	db.Table("area").Save(&area)
//	defer db.Close()
//	//删除一条数据
//	//db.Table("area").Delete(&area)
//}
//
///**
//自定义查询
//*/
//func main_() {
//	db, _ := gorm.Open("mysql", "root:123456@/test?charset=utf8&parseTime=True&loc=Local")
//	rows, _ := db.Raw("select id , area_name from area").Rows()
//	for rows.Next() {
//		var id int
//		var area_name string
//		rows.Scan(&id, &area_name)
//		fmt.Println(id, area_name)
//	}
//	defer db.Close()
//}
type User struct {
	Id        int       `json:"id" gorm:"column:id"`
	Name      string    `json:"name" gorm:"column:name"`
	Addr      string    `json:"addr" gorm:"column:addr"`
	Age       int       `json:"age" gorm:"column:age"`
	Birth     string    `json:"birth" gorm:"column:birth"`
	Sex       int       `json:"sex" gorm:"column:sex"`
	UpdatedAt time.Time `json:"update_at" gorm:"column:update_at"" description:"更新时间"`
	CreatedAt time.Time `json:"create_at" gorm:"column:create_at" description:"创建时间"`
}

/**
新建一张表
*/
func main() {
	DbHelper.AutoMigrate(&User{})
}
