package testServices

import (
	"fmt"

	"github.com/nickWoott/my-hex-api/internal/ports"
)

type TestService struct {
	DatabaseConnection ports.DatabaseConnection
}

func NewTestService(DatabaseConnection ports.DatabaseConnection) TestService {
	return TestService{
		DatabaseConnection: DatabaseConnection,
	}
}

func (ts *TestService) GetData(id string) (string, error) {
	storypoint, err := ts.DatabaseConnection.GetStoryPointById(id)

if err != nil {
	fmt.Println("there was an error")
	return storypoint, err
}
	// what does this return I am not sure 
	// call the function from the secondary adapter
	return storypoint, err

}
