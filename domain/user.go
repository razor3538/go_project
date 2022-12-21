package domain

// User model
type User struct {
	Base
	Email    string `gorm:"type:varchar(100)" json:"email"`
	Password string `gorm:"type:varchar(100)" json:"password"`
	Name     string `gorm:"type:varchar(100)" json:"name"`
	LastName string `gorm:"type:varchar(100)" json:"last_name"`
}
