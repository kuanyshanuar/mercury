package domain

// GenderType - specifies a particular sex type.
type GenderType string

// Supported question types.
const (
	GenderTypeFemale GenderType = "Female"
	GenderTypeMale   GenderType = "Male"
)

// Mapping of question types and their representation as numbers.
var (
	GenderTypeCodeMap = map[int]GenderType{
		1: GenderTypeFemale,
		2: GenderTypeMale,
	}
	GenderTypeValueMap = map[GenderType]int{
		GenderTypeFemale: 1,
		GenderTypeMale:   2,
	}
)

// IsValid checks if the provided question type is valid.
func (qt GenderType) IsValid() bool {
	_, ok := GenderTypeValueMap[qt]
	return len(qt) != 0 && ok
}

// String - string representation.
func (qt GenderType) String() string {
	return string(qt)
}
