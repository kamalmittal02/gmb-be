package injector

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	controller2 "github.com/kamalmittal01/girraj-sweet-showcase-BE/controller"
	"github.com/kamalmittal01/girraj-sweet-showcase-BE/dtos"
	"github.com/kamalmittal01/girraj-sweet-showcase-BE/repository/enquiry"
	service2 "github.com/kamalmittal01/girraj-sweet-showcase-BE/service"
	"time"
)

type Injector struct {
	EnquiryController controller2.EnquiryControllerI
}

func InitInjector(config *dtos.Config) *Injector {
	db := ConnectDB(config.Database.Url)
	repository := enquiry.NewEnquiryRepository(db)
	service := service2.NewEnquiryService(repository)
	controller := controller2.NewEnquiryController(service)
	return &Injector{
		EnquiryController: controller,
	}
}

func ConnectDB(databaseURl string) *pgxpool.Pool {
	if databaseURl == "" {
		fmt.Print("DATABASE_URL not set")
		panic("DATABASE_URL not set")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, databaseURl)
	if err != nil {
		fmt.Printf("failed to create connection pool: %v\n", err)
		panic("failed to create connection pool")
	}

	if err := pool.Ping(ctx); err != nil {
		fmt.Printf("failed to ping database: %v\n", err)
		panic("failed to ping database")
	}

	return pool
}
