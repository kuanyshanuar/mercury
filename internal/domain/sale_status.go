package domain

// SaleStatuses table name
const (
	SaleStatusesTableName = "sale_statuses"
)

// SaleStatusID - saleStatus id
type SaleStatusID int64

// SaleStatus - represents saleStatus
type SaleStatus struct {

	// ID - id of status.
	//
	ID SaleStatusID `json:"id" gorm:"primaryKey;autoIncrement;column:id"`

	// Name - name of status.
	//
	Name string `json:"name" gorm:"column:name"`
}
