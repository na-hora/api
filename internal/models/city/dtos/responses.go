package dtos

type ListCitiesByStateResponse struct {
	Cities []City `json:"cities"`
}

type City struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
