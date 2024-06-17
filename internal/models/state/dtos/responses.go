package dtos

type ListAllStatesResponse struct {
	States []State `json:"states"`
}

type State struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	UF   string `json:"uf"`
}
