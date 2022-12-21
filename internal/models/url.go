package models

// PathID struct for path binding
type PathID struct {
	ID string `uri:"id" json:"id" binding:"required"`
} //@name IDResponse
