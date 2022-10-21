package city

type City struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Region     string `json:"region"`
	District   int    `json:"district"`
	Foundation int    `json:"foundation"`
}
