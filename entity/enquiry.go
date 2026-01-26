package entity

import (
	"encoding/json"
	"time"
)

type Enquiry struct {
	Id        uint            `json:"id" gorm:"primaryKey;AutoIncrement"`
	Name      string          `json:"name"`
	Phone     string          `json:"phone"`
	Email     string          `json:"email"`
	Message   json.RawMessage `json:"message"`
	CreatedAt time.Time       `json:"created_at"`
}

type EnquiryResponse struct {
	Name      string          `json:"name"`
	Phone     string          `json:"phone"`
	Email     string          `json:"email"`
	Message   json.RawMessage `json:"message"`
	CreatedAt string          `json:"created_at"`
}

func (Enquiry) TableName() string {
	return "enquiry"
}
