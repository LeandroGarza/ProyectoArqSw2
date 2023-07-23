package dto

type ItemDto struct {
	Id        string  `json:"id"`
	Title     string  `json:"title"`
	UserId    int     `json:"userid"`
	Image     string  `json:"image"`
	Currency  string  `json:"currency"`
	Price     float32 `json:"price"`
	Sale_sate int     `json:"state"`
	Condition string  `json:"condition"`
	Address   string  `json:"address"`
}

type ItemsDto []ItemDto
