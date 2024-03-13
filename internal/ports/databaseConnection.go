package ports

import "github.com/nickWoott/my-hex-api/internal/core/domain"

type DatabaseConnection interface {
	GetStoryPointById(id int) (domain.StorypointRequest, error)
	PostStoryPoints(storypoint []domain.StorypointRequest) (error)
}

