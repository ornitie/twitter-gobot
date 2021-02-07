package repositories

type Server interface {
	Create() error
	Read() (interface{}, error)
	Update() (interface{}, error)
	Delete() (interface{}, error)
}
