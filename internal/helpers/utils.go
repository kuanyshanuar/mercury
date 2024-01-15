package helpers

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	errors "gitlab.com/zharzhanov/mercury/internal/error"

	"github.com/caarlos0/env"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

// UsageFor - print usage.
func UsageFor(fs *flag.FlagSet, short string) func() {
	return func() {
		fmt.Fprintf(os.Stderr, "USAGE\n")
		fmt.Fprintf(os.Stderr, "  %s\n", short)
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "FLAGS\n")
		w := tabwriter.NewWriter(os.Stderr, 0, 2, 2, ' ', 0)
		fs.VisitAll(func(f *flag.Flag) {
			fmt.Fprintf(w, "\t-%s %s\t%s\n", f.Name, f.DefValue, f.Usage)
		})
		w.Flush()
		fmt.Fprintf(os.Stderr, "\n")
	}
}

// ValidateConfig - validates configuration.
func ValidateConfig(config domain.AppConfig) error {

	err := validator.New().Struct(config)
	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		for _, err := range err.(validator.ValidationErrors) {
			return errors.NewErrFailedPrecondition(
				fmt.Sprintf("%s required", err.Field()),
			)
		}
	}
	// Validate supported environments.
	//
	if !domain.CanonicalEnvironments[config.Environment] {
		return errors.NewErrFailedPrecondition(
			fmt.Sprintf("Unsupported environment: %s", config.Environment),
		)
	}
	return nil
}

// LoadConfig - reads configuration from the environment variables.
func LoadConfig() (config domain.AppConfig, _ error) {
	if err := env.Parse(&config); err != nil {
		return config, fmt.Errorf("failed to read configuration: %v", err)
	}
	return config, nil
}

// LoadEmailConfig - reads configuration from the environment variables.
func LoadEmailConfig() (config domain.EmailConfig, _ error) {
	if err := env.Parse(&config); err != nil {
		return config, fmt.Errorf("failed to read configuration: %v", err)
	}
	return config, nil
}

// LoadEmailTemplate - reads configuration from the environment variables.
func LoadEmailTemplate() (config domain.EmailTemplate, _ error) {
	if err := env.Parse(&config); err != nil {
		return config, fmt.Errorf("failed to read configuration: %v", err)
	}
	return config, nil
}

// Paginate - gorm pagination.
func Paginate(page domain.PageRequest) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		db = db.Offset(page.Offset).Limit(page.Size)
		return db
	}
}

// PaginateRandom - gorm pagination with random offset.
func PaginateRandom(page domain.PageRequest) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := int(rand.Float64() * float64(page.Offset))
		db = db.Offset(offset).Limit(page.Size)
		return db
	}
}

// PrepareOrder - prepare order for a query
func PrepareOrder(sorts []domain.Sort) string {
	var order []string
	for _, sort := range sorts {
		order = append(order,
			fmt.Sprintf(
				`"%s" %s`,
				sort.FieldName,
				sort.Order,
			),
		)
	}
	return strings.Join(order, ", ")
}

// GetTotal - get total of data in a storage by table name
func GetTotal(db gorm.DB, tableName string) (int64, error) {
	var totalCount int64
	err := db.
		Table(tableName).
		Where("deleted_at = 0").
		Count(&totalCount).Error
	if err != nil {
		return 0, err
	}

	return totalCount, nil
}

// StartOfTheDay - returns start of the day
func StartOfTheDay(timestamp int64) int64 {
	return time.Unix(timestamp, 0).Truncate(24 * time.Hour).Unix()
}

// EndOfTheDay - returns end of the day
func EndOfTheDay(timestamp int64) int64 {
	return time.Unix(timestamp, 0).Truncate(24 * time.Hour).Add(24 * time.Hour).Unix()
}
