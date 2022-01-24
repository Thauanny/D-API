package entities

type Food struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       [2]string `json:"image"`
	Price       float32   `json:"price"`
}
