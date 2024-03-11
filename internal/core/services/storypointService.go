package services

import (
	"fmt"

	"github.com/nickWoott/my-hex-api/internal/core/domain"
	"github.com/nickWoott/my-hex-api/internal/ports"
)

type StoryPointService struct {
	DatabaseConnection ports.DatabaseConnection
}

func NewStoryPointService(DatabaseConnection ports.DatabaseConnection) StoryPointService {
	return StoryPointService{
		DatabaseConnection: DatabaseConnection,
	}
}

func (ss *StoryPointService) GetData(id string) (domain.StorypointRequest, error) {
	storypoint, err := ss.DatabaseConnection.GetStoryPointById(id)

if err != nil {
	fmt.Println("there was an error in the service")
	return storypoint, err
}
	return storypoint, err

}

//this may not be true hexagonal architecture, as it uses an interface from the input
func(ss *StoryPointService) SendStoryPoints(storypoint *domain.StorypointRequest ) (error) {

 err := ss.DatabaseConnection.PostStoryPoint(*storypoint) 
 
 if err != nil {
fmt.Println("error in the service")
return err
	}	

	return nil
	
}