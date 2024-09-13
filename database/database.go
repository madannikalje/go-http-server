package database

import (
	"http-test/config"
	"http-test/models"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

var DB *pg.DB

func InitializeDatabase() {
	DB = pg.Connect(&pg.Options{
		Addr:     config.DB_HOST + ":" + config.DB_PORT,
		User:     config.DB_USER,
		Password: config.DB_PASSWORD,
		Database: config.DB_NAME,
	})

	// err := createSchema(DB)
	// if err != nil {
	// 	panic(err)
	// }
}

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{(*models.User)(nil), (*models.Post)(nil)} {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp: false,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
