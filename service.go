package main

import "context"

type AuthorRepo interface {
	Create(context.Context, Author) error 
}

type Service struct {
	authorRepo AuthorRepo
}

func NewService(authorRepo AuthorRepo) *Service {
	return &Service{
		authorRepo: authorRepo,
	}
}

func (s *Service) CreateAuthor(ctx context.Context, author Author) error {
	return s.authorRepo.Create(ctx, author)
}
