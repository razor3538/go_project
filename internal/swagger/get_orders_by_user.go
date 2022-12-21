package swagger

import "github.com/gofrs/uuid"

// GetOrdersByUser user model for /api/user/orders request route
type GetOrdersByUser struct {
	UserID uuid.UUID `json:"user_id" example:""`
} //@name GetOrdersByUser
