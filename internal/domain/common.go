package domain

// Response statuses.
const (
	ResponseStatusSuccess = ResponseStatus("success")
	ResponseStatusError   = ResponseStatus("error")
)

// ResponseStatus - represents the response status.
type ResponseStatus string

// Bodier may be implemented by the gateway response types.
// It's not necessary for your response types to implement Bodier, but it may
// help for more sophisticated use cases.
type Bodier interface {
	GetBody() interface{}
}

// Default values.
const (
	DefaultPageNumber = 0 // page starts from "0"
	DefaultPageSize   = 25
	//DefaultIncludeDeleted = false
)

// PageRequest - request the page.
type PageRequest struct {
	Offset int // page offset (starts from 0)
	Size   int // page size
}

// NextPage - move request to next page.
func (pr *PageRequest) NextPage() {
	pr.Offset++
}

// Total - represents the total number of records in the data source.
type Total int64

const (
	// Asc - the ascending order.
	Asc = "ASC"
	// Desc - the descending order.
	Desc = "DESC"
)

// Sort - the structure holds the sorting and ordering parameters.
type Sort struct {
	FieldName string // the name of the field to sort
	Order     string // order by (asc or desc)
}
