package migration

import (
	"github.com/isnakolah/task-tracker-backend/model"
	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Task{})
	db.AutoMigrate(&model.TaskGroup{})
}
