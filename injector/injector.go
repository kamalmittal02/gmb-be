package injector

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	controller2 "github.com/kamalmittal01/girraj-sweet-showcase-BE/controller"
	service2 "github.com/kamalmittal01/girraj-sweet-showcase-BE/service"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
	"os"
	"strings"
)

type Injector struct {
	EnquiryController controller2.EnquiryControllerI
}

func InitInjector() *Injector {
	//db := ConnectDB(config.Database.Url)
	rctx := context.Background()
	//repository := enquiry.NewEnquiryRepository(db)
	sheetClient, err := SheetsService(rctx)
	if err != nil {
		fmt.Printf("error initializing sheets service: %v\n", err)
		panic("error initializing sheets service")
	}
	sheetsService := service2.NewSheetsService(sheetClient)
	service := service2.NewEnquiryService(sheetsService)

	controller := controller2.NewEnquiryController(service)
	return &Injector{
		EnquiryController: controller,
	}
}

func ConnectDB(databaseURl string) *pgxpool.Pool {
	//if databaseURl == "" {
	//	fmt.Print("DATABASE_URL not set")
	//	panic("DATABASE_URL not set")
	//}
	//
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//
	//pool, err := pgxpool.New(ctx, databaseURl)
	//if err != nil {
	//	fmt.Printf("failed to create connection pool: %v\n", err)
	//	panic("failed to create connection pool")
	//}
	//
	//if err := pool.Ping(ctx); err != nil {
	//	fmt.Printf("failed to ping database: %v\n", err)
	//	panic("failed to ping database")
	//}
	//
	return nil
}

func SheetsService(ctx context.Context) (*sheets.Service, error) {
	val := os.Getenv("GOOGLE_CREDENTIALS_JSON")
	if val == "" {
		return nil, fmt.Errorf("GOOGLE_CREDENTIALS_JSON not set")
	}

	var creds []byte
	var err error

	// If env contains JSON directly (Render)
	if strings.HasPrefix(strings.TrimSpace(val), "{") {
		creds = []byte(val)
	} else {
		// Otherwise treat as file path (local)
		creds, err = os.ReadFile(val)
		if err != nil {
			return nil, fmt.Errorf("failed reading credentials file: %w", err)
		}
	}

	config, err := google.JWTConfigFromJSON(
		creds,
		sheets.SpreadsheetsScope,
	)
	if err != nil {
		return nil, fmt.Errorf("invalid google credentials: %w", err)
	}

	client := config.Client(ctx)
	return sheets.New(client)
}
