package conf

import (
	model "beego-basic/models"
	"fmt"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// DB
var Db *gorm.DB

// DsetUpda
func SetUpConnectDb() {
	dsn := "sqlserver://sa:Quanit04042001@192.168.9.22:1433?database=ThucTap"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err == nil {
		Db = db
		modelMirate := []interface{}{}
		modelMirate = append(modelMirate, model.InitUser()...)
		for _, value := range modelMirate {
			Db.AutoMigrate(value)
		}

	} else {
		fmt.Println(err)
	}
}
