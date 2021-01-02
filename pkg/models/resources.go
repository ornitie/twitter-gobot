package models

type (
	ResourceError struct {
		Message string
	}
)

func (resourceError ResourceError) Error() string {
	return resourceError.Message
}
