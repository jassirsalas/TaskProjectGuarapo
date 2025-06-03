package models

type Task struct {
	ID        int    `gorm:"primaryKey; autoIncrement" json:"id"`
	Title     string `gorm:"notnull" json:"title" binding:"required"`
	Completed bool   `gorm:"default:false" json:"completed"`
	Owner     string `gorm:"type:varchar(64)" json:"owner"` // el username dueño de la tarea
}

type LoginRequest struct {
	Username string `json:"username"`
}
