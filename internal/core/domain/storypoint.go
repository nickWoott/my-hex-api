package domain

type Choices struct {
	Destination string
	Choice string
	}
	
	
	type StorypointRequest  struct {
		Id string
		Text string 
		Choices []Choices
		
	}
	