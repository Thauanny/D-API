package entities

type Order struct {
	Number string `json:"number"`
	Type   string `json:"type"`
	Items  []Item `json:"items"`
}
