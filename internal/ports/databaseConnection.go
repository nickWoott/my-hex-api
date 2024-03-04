package ports

type DatabaseConnection interface {
	GetStoryPointById(id string) string
}
