package migration

import (
	"os/user"

	"github.com/isnakolah/task-tracker-backend/task"
	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&task.Task{})
	db.AutoMigrate(&task.TaskGroup{})
}
