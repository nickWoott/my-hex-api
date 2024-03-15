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

func (ss *StoryPointService) GetData(id int) (domain.StorypointRequest, error) {
	storypoint, err := ss.DatabaseConnection.GetStoryPointById(id)

if err != nil {
	fmt.Println("there was an error in the service")
	return storypoint, err
}
	return storypoint, err

}

//this may not be true hexagonal architecture, as it uses an interface from the input
func(ss *StoryPointService) SendStoryPoints(storyPoints []domain.StorypointRequest ) (error) {

 err := ss.DatabaseConnection.PostStoryPoints(storyPoints) 
 
 if err != nil {
fmt.Println("error in the service")
return err
	}	

	return nil
	
}

func(ss *StoryPointService) DeleteAll() (error) {

	err := ss.DatabaseConnection.DeleteAllStoryPoints()

	if err != nil{
		return err
	}

	return nil
}