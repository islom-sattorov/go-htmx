package services

import (
	"context"

	"go_role/internal/core/domain"
	"go_role/internal/core/ports"
)

type PostService struct {
	postRepo ports.PostRepository
	userRepo ports.UserRepository
}

func NewPostService(pr ports.PostRepository, ur ports.UserRepository) *PostService {
	return &PostService{
		postRepo: pr,
		userRepo: ur,
	}
}

func (s *PostService) CreatePost(ctx context.Context, create *domain.PostCreate) (*domain.Post, error) {
	// Implementation
}

func (s *PostService) GetPost(ctx context.Context, id uint) (*domain.Post, error) {
	// Implementation
}
