package swagger

// CreateUser struct
type CreateUser struct {
	Email    string `gorm:"type:varchar(100);unique;not null" json:"email" example:"admin@mail.ru" binding:"required"`
	Name     string `gorm:"type:varchar(100);not null" json:"name" example:"Eduard" binding:"required"`
	LastName string `gorm:"type:varchar(100);not null" json:"last_name" example:"Gaifutdinov" binding:"required"`
	Password string `gorm:"type:varchar(100);not null" json:"password" example:"123" binding:"required"`
}
