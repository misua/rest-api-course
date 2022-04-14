package comment

import (
	"context"
	"fmt"
)

// comment - a representation of the comment
//structure for our service
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

// Service - is the struct on which all our logic
// will be built on top of
type Service struct {
	Store Store
}

//NewService - returns a pointer to new
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

//constructor func
func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("retrieving a comment")
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, err
	}
	return cmt, nil
}
