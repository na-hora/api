package dtos

type ListCitiesByStateResponse struct {
	Cities []City `json:"cities"`
}

type GetCityByIBGEResponse struct {
	City City `json:"city"`
}

type City struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
