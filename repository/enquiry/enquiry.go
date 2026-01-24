package enquiry

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	sqlc "github.com/kamalmittal01/girraj-sweet-showcase-BE/repository/enquiry/postgres/sqlc"

	"github.com/kamalmittal01/girraj-sweet-showcase-BE/entity"
)

type EnquiryRepositoryI interface {
	Create(ctx context.Context, e entity.Enquiry) error
	GetAll(ctx context.Context) ([]entity.Enquiry, error)
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

func (r *EnquiryRepository) GetAll(ctx context.Context) ([]entity.Enquiry, error) {
	//rows, err := r.db.Query(ctx, `
	//	SELECT id, name, phone, email, query, created_at
	//	FROM enquiries
	//	ORDER BY created_at DESC
	//`)
	//if err != nil {
	//	return nil, err
	//}
	//defer rows.Close()
	//
	//var enquiries []entity.Enquiry
	//for rows.Next() {
	//	var e entity.Enquiry
	//	if err := rows.Scan(
	//		&e.Id,
	//		&e.Name,
	//		&e.Phone,
	//		&e.Email,
	//		&e.Message,
	//		&e.CreatedAt,
	//	); err != nil {
	//		return nil, err
	//	}
	//	enquiries = append(enquiries, e)
	//}

	//return enquiries, nil
	return nil, nil
}
