package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kamalmittal01/girraj-sweet-showcase-BE/entity"
	"github.com/kamalmittal01/girraj-sweet-showcase-BE/request"
)

type EnquiryServiceI interface {
	CreateEnquiry(ctx context.Context, enquiry request.Enquiry) error
}

type EnquiryService struct {
	SheetService SheetsServiceI
}

func NewEnquiryService(sheetService SheetsServiceI) EnquiryServiceI {
	return &EnquiryService{SheetService: sheetService}
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
	//err = es.EnquiryRepo.Create(ctx, enquiryEntity)
	//if err != nil {
	//	return err
	//}

	// Append to Google Sheets
	err = es.SheetService.AppendEnquiryToSheet(ctx, es.SheetService.(*SheetsService).sheetsClient, enquiryEntity)
	if err != nil {
		fmt.Printf("error appending enquiry to sheet: %v\n", err)
	}
	return nil
}

//func (es *EnquiryService) GetEnquiry(ctx context.Context, createdAt time.Time) (*[]entity.Enquiry, error) {
//	res, err := es.EnquiryRepo.GetAll(ctx, createdAt)
//	if err != nil {
//		fmt.Println(fmt.Errorf("error getting enquiries: %v", err))
//		return nil, err
//	}
//	return res, nil
//}
