package testServices

type TestService struct {
}

func NewTestService() TestService {
	return TestService{}
}

func (ts *TestService) GetData(id string) string {
	return id
}
