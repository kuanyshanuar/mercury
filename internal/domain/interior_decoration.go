package domain

// InteriorDecorationID - id of interior decoration.
type InteriorDecorationID int64

// InteriorDecoration - represents interior decoration type.
type InteriorDecoration struct {
	// ID - id of interior decoration type
	//
	ID InteriorDecorationID `json:"id" gorm:"primaryKey;autoIncrement;column:id"`

	// Name - name of interior decoration type.
	//
	Name string `json:"name" gorm:"not null;column:name"`
}
