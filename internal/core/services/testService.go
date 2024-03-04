package testServices

import "github.com/nickWoott/my-hex-api/internal/ports"

type TestService struct {
	DatabaseConnection ports.DatabaseConnection
}

func NewTestService(DatabaseConnection ports.DatabaseConnection) TestService {
	return TestService{
		DatabaseConnection: DatabaseConnection,
	}
}

func (ts *TestService) GetData(id string) string {
	storypoint := ts.DatabaseConnection.GetStoryPointById(id)

	// what does this retu
	//call the function from the secondary adapter
	return storypoint

}
