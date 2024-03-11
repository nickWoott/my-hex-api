package ports

import "github.com/nickWoott/my-hex-api/internal/core/domain"

type DatabaseConnection interface {
	GetStoryPointById(id string) (domain.StorypointRequest, error)
	PostStoryPoint(storypoint domain.StorypointRequest) (error)
}

