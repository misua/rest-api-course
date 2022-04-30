package comment

import (
	"context"
	"errors"
	"fmt"
)

//custom error handling msg
var (
	ErrFetchingComment = errors.New("failed to fetch comment by id")
	ErrNotImplemented  = errors.New("not implemented")
)

// comment - a representation of the comment
//structure for our service
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

//Store - this interface defines all of the methods
// that our service needs in order to operate

type Store interface {
	GetComment(context.Context, string) (Comment, error)
	PostComment(context.Context, Comment) (Comment, error)
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
		return Comment{}, ErrFetchingComment
	}
	return cmt, nil
}

func (s *Service) UpdateComment(ctx context.Context, cmt Comment) error {
	return ErrNotImplemented
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return ErrNotImplemented
}

func (s *Service) PostComment(ctx context.Context, cmt Comment) (Comment, error) {
	insertCmt, err := s.Store.PostComment(ctx, cmt)

	if err != nil {
		return Comment{}, err
	}

	return insertCmt, nil
}
