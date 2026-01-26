package enquiry

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kamalmittal01/girraj-sweet-showcase-BE/entity"
	sqlc "github.com/kamalmittal01/girraj-sweet-showcase-BE/repository/enquiry/postgres/sqlc"
	"time"
)

type EnquiryRepositoryI interface {
	Create(ctx context.Context, e entity.Enquiry) error
	GetAll(ctx context.Context, createdAt time.Time) (*[]entity.Enquiry, error)
}

type EnquiryRepository struct {
	reader *sqlc.Queries
}

func NewEnquiryRepository(db *pgxpool.Pool) EnquiryRepositoryI {
	return &EnquiryRepository{
		reader: sqlc.New(db),
	}
}

func (r *EnquiryRepository) Create(ctx context.Context, entity entity.Enquiry) error {
	var email pgtype.Text
	if entity.Email != "" {
		email = pgtype.Text{
			String: entity.Email,
			Valid:  true,
		}
	}
	err := r.reader.CreateEnquiry(ctx, sqlc.CreateEnquiryParams{
		Name:    entity.Name,
		Phone:   entity.Phone,
		Email:   email,
		Message: entity.Message,
	})
	if err != nil {
		fmt.Printf("error creating enquiry: %v\n", err)
		return err
	}
	return nil
}

func (r *EnquiryRepository) GetAll(ctx context.Context, createdAt time.Time) (*[]entity.Enquiry, error) {
	rows, err := r.reader.GetEnquiries(ctx, pgtype.Timestamptz{
		Time:  createdAt,
		Valid: true,
	})
	if err != nil {
		return nil, err
	}

	var enquiries []entity.Enquiry
	for _, row := range rows {
		var email string
		if row.Email.Valid {
			email = row.Email.String
		}
		enquiries = append(enquiries, entity.Enquiry{
			Id:        uint(row.ID),
			Name:      row.Name,
			Phone:     row.Phone,
			Email:     email,
			Message:   row.Message,
			CreatedAt: row.CreatedAt.Time,
		})
	}

	//return enquiries, nil
	return &enquiries, nil
}
