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

func MigrateUp() error {
	models := getModel()

	for _, value := range models {
		if err := App.App.DB.AutoMigrate(value); err != nil {
			return err
		}
	}
	return nil
}

func MigrateDown() error {
	models := getModel()
	for _, tableName := range models {
		if err := App.App.DB.Migrator().DropTable(tableName); err != nil {
			return err
		}
	}
	return nil
}

func MigrateDownTable(table interface{}) error {
	if err := App.App.DB.Migrator().DropTable(table); err != nil {
		return err
	}
	return nil
}
