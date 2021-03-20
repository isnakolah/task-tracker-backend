package account

import (
	"github.com/isnakolah/task-tracker-backend/task"
	"github.com/isnakolah/task-tracker-backend/utils"
)

type User struct {
	utils.Base
	Email      string           `json:"email"`
	FirstName  string           `json:"first_name"`
	LastName   string           `json:"last_name"`
	Password   string           `json:"password"`
	Tasks      []task.Task      `json:"tasks"`
	TaskGroups []task.TaskGroup `json:"task_groups"`
}
