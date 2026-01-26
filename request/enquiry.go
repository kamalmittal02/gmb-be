package request

import "time"

type Enquiry struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Message string `json:"message"`
}

type EnquiryFilter struct {
	CreatedAt time.Time `json:"created_at" binding:"required"`
}
