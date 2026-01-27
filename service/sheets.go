package service

import (
	"context"
	"github.com/kamalmittal01/girraj-sweet-showcase-BE/entity"
	"google.golang.org/api/sheets/v4"
	"os"
	"time"
)

type SheetsServiceI interface {
	AppendEnquiryToSheet(ctx context.Context, svc *sheets.Service, enquiry entity.Enquiry) error
}

type SheetsService struct {
	sheetsClient *sheets.Service
}

func NewSheetsService(client *sheets.Service) SheetsServiceI {
	return &SheetsService{sheetsClient: client}
}

func (s *SheetsService) AppendEnquiryToSheet(ctx context.Context, svc *sheets.Service, enquiry entity.Enquiry) error {
	values := []interface{}{
		enquiry.Name,
		enquiry.Phone,
		enquiry.Email,
		string(enquiry.Message),
		time.Now().Format(time.RFC3339),
	}

	vr := &sheets.ValueRange{
		Values: [][]interface{}{values},
	}

	_, err := svc.Spreadsheets.Values.Append(
		os.Getenv("GOOGLE_SHEET_ID"),
		"GMB_Sweets_Enquiry!A:F",
		vr,
	).
		ValueInputOption("RAW").
		InsertDataOption("INSERT_ROWS").
		Do()

	return err
}
