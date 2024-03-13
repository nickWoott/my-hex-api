package domain

type Choices struct {
	Destination string `json:"destination"`
	Choice string `json:"choice"`
	}
	
	
	type StorypointRequest  struct {
		Id string `json:"id"`
		Text string `json:"text"`
		Choices []Choices `json:"choices"`
		
	}
	