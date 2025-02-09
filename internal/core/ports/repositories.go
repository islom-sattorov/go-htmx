package ports

import (
	"context"

	"go_role/internal/core/domain"
)

type PostRepository interface {
	Create(ctx context.Context, post *domain.Post) error
	GetByID(ctx context.Context, id uint) (*domain.Post, error)
	List(ctx context.Context, page, limit int) ([]domain.Post, error)
	Update(ctx context.Context, post *domain.Post) error
	Delete(ctx context.Context, id uint) error
}

type UserRepository interface {
	// User repository methods
}
