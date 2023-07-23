package dtos

type ItemSolrDto struct {
	Id        string    `json:"id"`
	Title     []string  `json:"title"`
	Userid    []int     `json:"userid"`
	Image     []string  `json:"image"`
	Currency  []string  `json:"currency"`
	Price     []float32 `json:"price"`
	Sale_sate []int     `json:"state"`
	Condition []string  `json:"condition"`
	Address   []string  `json:"address"`
	Version   int64     `json:"_version_"`
}

type ItemsSolrDto []ItemSolrDto
