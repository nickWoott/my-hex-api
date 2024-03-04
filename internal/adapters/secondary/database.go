package database

import (
	"fmt"
	"os"
)

type LocalDatabaseConnection struct {
}

func NewLocalDatabase() LocalDatabaseConnection {
	return LocalDatabaseConnection{}
}

func (ld *LocalDatabaseConnection) GetStoryPointById(id string) string {

	fileContents, err := os.ReadFile("/Users/welldunn/Code-Projects/go-projects/my-hex-api/data/data.json")

	if err != nil {
		fmt.Print("THERE IS AN ERROR", err)
	}

	return string(fileContents)
}
