package service

import (
	"Cloudflare-Assistant/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	Conf = &model.Config{}
	DB   = &gorm.DB{}
)

func InitService() {
	InitConfigFromPath("config/config.yaml")
	InitDBFromPath("config/db.sqlite")
}

// InitConfigFromPath init config from path
func InitConfigFromPath(path string) {
	err := Conf.Read(path)
	if err != nil {
		panic(err)
	}
}

// InitDBFromPath init db from path
func InitDBFromPath(path string) {
	var err error
	DB, err = gorm.Open(sqlite.Open(path), &gorm.Config{
		CreateBatchSize: 200,
	})
	if err != nil {
		panic(err)
	}
	if Conf.Debug {
		DB.Debug()
	}
	err = DB.AutoMigrate(
		&model.User{},
	)

}
