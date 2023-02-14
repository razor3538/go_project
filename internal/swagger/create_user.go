package swagger

// CreateUser struct
type CreateUser struct {
	Login    string `gorm:"type:varchar(100);unique;not null" json:"login" example:"admin" binding:"required"`
	Password string `gorm:"type:varchar(100);not null" json:"password" example:"123" binding:"required"`
}
