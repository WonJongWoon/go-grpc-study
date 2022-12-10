package todo

import "gorm.io/gorm"

type Status string

const (
	Backlog    Status = "backlog"
	Accepted   Status = "accepted"
	InProgress Status = "in_progress"
	Occupied   Status = "occupied"
)

type Todo struct {
	gorm.Model
	Title       string `gorm:"column:title"`
	Description string `gorm:"column:description"`
	Status      Status `gorm:"column:status"`
	Author      string `gorm:"column:author"`
}

func New(title string, description string, author string) *Todo {
	return &Todo{
		Title:       title,
		Description: description,
		Status:      Backlog,
		Author:      author,
	}
}
