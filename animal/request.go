package animal

type AnimalRequest struct {
	Name  string `json:"name" binding:"required"`
	Class string `json:"class" binding:"required"`
	Legs  int    `json:"legs" binding:"number"`
}

type UpdateAnimalRequest struct {
	Name  string `json:"name"`
	Class string `json:"class"`
	Legs  int    `json:"legs"`
}
