package domain

// FilterKeysTableName - table name
const (
	FilterKeysTableName = "filter_keys"
)

// FilterKey - filter key
type FilterKey struct {
	// ID - id of parking type.
	//
	ID int64 `json:"id" gorm:"primaryKey;autoIncrement;column:id"`

	// Key - name of parking type.
	//
	Key string `json:"key" gorm:"not null;column:key"`
}
