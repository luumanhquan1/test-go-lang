package models

import (
	"time"
)

func InitUser() []interface{} {
	result := []interface{}{SoThich{}, Img{}, User{}}
	return result
}

type User struct {
	Id               string    `json:"id" gorm:"primary_key"`
	Msv              string    `json:"msv" gorm:"size(50)"`
	Name             string    `json:"name" gorm:"size(50)"`
	GioiThieuBanThan string    `json:"gioiThieuBanThan"`
	Lop              string    `json:"lop"`
	GioiTinh         bool      `json:"gioiTinh"`
	ChucDanh         string    `json:"chucDanh"`
	Address          string    `json:"address"`
	NgaySinh         time.Time `json:"ngaySinh"`
	SoThich          []SoThich `json:"soThichs" gorm:"foreignKey:UserId;references:Id"`
	Img              []Img     `json:"images" gorm:"foreignKey:UserId;references:Id"`
}

type SoThich struct {
	Id     string `json:"id" gorm:"primary_key"`
	Name   string `json:"name"`
	UserId string `json:"userId"`
}

type Img struct {
	Id     string `json:"id" gorm:"primary_key"`
	Url    string `json:"url"`
	UserId string `json:"userId"`
}
