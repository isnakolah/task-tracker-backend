package task

import (
	"time"

	"github.com/isnakolah/task-tracker-backend/utils"
)

type Task struct {
	utils.Base
	Name        string    `json:"name"`
	Description string    `json:"description"`
	SetReminder bool      `json:"set_reminder"`
	DueOn       time.Time `json:"due_on"`
	UserID      uint      `json:"user_id"`
	TaskGroupID uint      `json:"task_group_id"`
}

type TaskGroup struct {
	utils.Base
	UserID uint   `json:"user_id"`
	Tasks  []Task `json:"tasks"`
}
