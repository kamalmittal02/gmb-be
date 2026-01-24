package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kamalmittal01/girraj-sweet-showcase-BE/entity"
	"github.com/kamalmittal01/girraj-sweet-showcase-BE/repository/enquiry"
	"github.com/kamalmittal01/girraj-sweet-showcase-BE/request"
)

type EnquiryServiceI interface {
	CreateEnquiry(ctx context.Context, enquiry request.Enquiry) error
}

type EnquiryService struct {
	EnquiryRepo enquiry.EnquiryRepositoryI
}

func NewEnquiryService(repo enquiry.EnquiryRepositoryI) EnquiryServiceI {
	return &EnquiryService{EnquiryRepo: repo}
}
func (es *EnquiryService) CreateEnquiry(ctx context.Context, enquiry request.Enquiry) error {
	message, err := json.Marshal(enquiry.Message)
	if err != nil {
		fmt.Printf("error marshalling message: %v\n", err)
	}
	enquiryEntity := entity.Enquiry{
		Name:    enquiry.Name,
		Phone:   enquiry.Phone,
		Email:   enquiry.Email,
		Message: message,
	}

	// Call the repository's Create method
	err = es.EnquiryRepo.Create(ctx, enquiryEntity)
	if err != nil {
		return err
	}

	return nil
}
