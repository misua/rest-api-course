package comment

// comment - a representation of the comment
//structure for our service
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

type Service struct{}

func NewService() *Service {
	return &Service{}
}
