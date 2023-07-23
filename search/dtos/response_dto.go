package dtos

type ResponseDto struct {
	ResponseHeader struct {
		Status int `json:"status"`
		QTime  int `json:"QTime"`
		Params struct {
			Q string `json:"q"`
		} `json:"params"`
	} `json:"responseHeader"`
	Response struct {
		NumFound      int          `json:"numFound"`
		Start         int          `json:"start"`
		NumFoundExact bool         `json:"numFoundExact"`
		Docs          ItemsSolrDto `json:"docs"`
	} `json:"response"`
}
