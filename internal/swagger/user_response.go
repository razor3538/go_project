package swagger

import "github.com/gofrs/uuid"

// UserResponse struct
type UserResponse struct {
	ID       uuid.UUID `gorm:"type:uuid;" json:"id"`
	Name     string    `gorm:"type:varchar(100);" json:"name"`
	LastName string    `gorm:"type:varchar(100);" json:"last_name"`
	Email    string    `gorm:"type:varchar(100);unique;not null" json:"email"`
}
