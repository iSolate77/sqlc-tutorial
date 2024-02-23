package core

import "context"

type AuthorRepo interface {
	Create(ctx context.Context, auth Author) (Author, error)
}

type Service struct {
	authorRepo AuthorRepo
}

func NewService(authorRepo AuthorRepo) *Service {
	return &Service{
		authorRepo: authorRepo,
	}
}

func (s *Service) CreateAuthor(ctx context.Context, auth Author) (Author, error) {
	return s.authorRepo.Create(ctx, auth)
}
