package domain

type StoryPoint struct {
	ID        string   `json:"id" bson:"id"`
	Story     string   `json:"storypoint" bson:"storypoint"`
	Choices   []string `json:"choices" bson:"choices"`
}

