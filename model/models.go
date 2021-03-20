package model

import "time"

type User struct {
	Base
	Email      string      `json:"email"`
	FirstName  string      `json:"first_name"`
	LastName   string      `json:"last_name"`
	Password   string      `json:"password"`
	Tasks      []Task      `json:"tasks"`
	TaskGroups []TaskGroup `json:"task_groups"`
}

type Task struct {
	Base
	Name        string    `json:"name"`
	Description string    `json:"description"`
	SetReminder bool      `json:"set_reminder"`
	DueOn       time.Time `json:"due_on"`
	UserID      uint      `json:"user_id"`
	TaskGroupID uint      `json:"task_group_id"`
}

type TaskGroup struct {
	Base
	UserID uint   `json:"user_id"`
	Tasks  []Task `json:"tasks"`
}

type Base struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
