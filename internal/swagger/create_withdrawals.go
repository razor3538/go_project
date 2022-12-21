package swagger

// CreateWithdrawals struct
type CreateWithdrawals struct {
	Order string  `json:"order" example:"2377225624" binding:"required"`
	Sum   float64 `json:"sum" example:"500" binding:"required"`
}
