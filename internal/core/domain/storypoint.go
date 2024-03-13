package domain

type Choices struct {
	Destination int `json:"destination"`
	Choice string `json:"choice"`
	}
	
	
	type StorypointRequest  struct {
		Id int `json:"id"`
		Text string `json:"text"`
		Choices []Choices `json:"choices"`
		
	}
	