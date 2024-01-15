package domain

// PurchaseMethodsTableName - table name
const (
	PurchaseMethodsTableName = "purchase_methods"
)

// PurchaseMethodID - purchase method id.
type PurchaseMethodID int64

// PurchaseMethod - represents purchase method struct
type PurchaseMethod struct {
	// ID - id of the purchase method.
	//
	ID PurchaseMethodID `json:"id" gorm:"primaryKey;autoIncrement;column:id"`

	// Name - name of purchase method.
	//
	Name string `json:"name" gorm:"not null;column:name"`
}
