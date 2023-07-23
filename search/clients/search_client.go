package clients

/*
import (
	"fmt"
	"net/http"
	"net/url"
)

func Search(query string) (*http.Response, error) {
	baseurl := "http://localhost:8983/solr/items/query"
	// instanciamos url.Values y lo establecemos con el valor de busqueda
	params := url.Values{}
	params.Set("q", "title:*"+query+"* OR condition:*"+query+"* OR address:*"+query+"*")

	// codificar los parametros
	url := baseurl + "?" + params.Encode()

	// solicitud http get a solr
	r, err := http.Get(url)
	if err != nil {
		return r, err
	}
	fmt.Printf("url al que le hacemos la peticion: %+v\n", url)
	fmt.Printf("Body del response: %+v\n", r.Body)
	return r, nil
}

*/
