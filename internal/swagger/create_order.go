package swagger

// CreateOrder struct
type CreateOrder struct {
	Number string `gorm:"type:varchar(100);not null" json:"number" example:"12345678903" binding:"required"`
}
