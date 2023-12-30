package Migrator

import (
	"github.com/amirhossein2831/DummyGo/src/App"
	"github.com/amirhossein2831/DummyGo/src/Model"
)

func getModel() (models []interface{}) {
	models = append(models,
		//add Model to Migrate here
		&Model.User{},
	)
	return
}

func Migrate() error {
	models := getModel()

	for _, value := range models {
		if err := App.App.DB.AutoMigrate(value); err != nil {
			return err
		}
	}
	return nil
}
