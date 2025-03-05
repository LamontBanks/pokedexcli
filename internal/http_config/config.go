package httpconfig

// Struct to save info between HTTP calls, ex: pagination variables
type Config struct {
	NextUrl     *string
	PreviousUrl *string
}
