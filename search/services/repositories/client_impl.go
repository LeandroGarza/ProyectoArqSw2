package repositories

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	log "github.com/sirupsen/logrus"
)

type SearchClient struct{}

func NewSearchClient() *SearchClient {
	return &SearchClient{}
}

func (sc *SearchClient) Search(query string) (*http.Response, error) {
	baseurl := "http://localhost:8983/solr/items/query"
	// instanciamos url.Values y lo establecemos con el valor de busqueda
	params := url.Values{}
	params.Set("q", "title:*"+query+"* OR condition:*"+query+"* OR address:*"+query+"* OR id:*"+query+"*")

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

func (sc *SearchClient) SearchByUserId(id int) (*http.Response, error) {
	baseurl := "http://localhost:8983/solr/items/query"

	params := url.Values{}
	params.Set("q", "userid:"+fmt.Sprintf("%v", id))

	url := baseurl + "?" + params.Encode()

	r, err := http.Get(url)
	if err != nil {
		log.Error()
		return r, err
	}
	fmt.Printf("url al que le hacemos la peticion: %+v\n", url)
	fmt.Printf("Body del response: %+v\n", r.Body)
	return r, nil
}

func (sc *SearchClient) DeleteByUserId(userid int) (*http.Response, error) {
	// url base
	baseurl := "http://localhost:8983/solr/items/update"

	// URL completa: http://localhost:8983/solr/items/update?commit=true
	url := fmt.Sprintf("%s?commit=true", baseurl)

	// Crear el cuerpo XML de la consulta de eliminación
	query := fmt.Sprintf("<delete><query>userid:%d</query></delete>", userid)

	// Realizar la solicitud POST a la URL
	resp, err := http.Post(url, "application/xml", strings.NewReader(query))
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()

	fmt.Printf("url al que le hacemos la peticion: %+v\n", url)
	fmt.Printf("Body del response: %+v\n", resp.Body)
	return resp, nil
}

func (sc *SearchClient) DeleteAll() (*http.Response, error) {
	// url base
	baseurl := "http://localhost:8983/solr/items/update"

	// URL completa: http://localhost:8983/solr/items/update?commit=true
	url := fmt.Sprintf("%s?commit=true", baseurl)

	// Crear el cuerpo XML de la consulta de eliminación
	query := "<delete><query>*:*</query></delete>"

	// Realizar la solicitud POST a la URL
	resp, err := http.Post(url, "application/xml", strings.NewReader(query))
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()

	fmt.Printf("url al que le hacemos la peticion: %+v\n", url)
	fmt.Printf("Body del response: %+v\n", resp.Body)
	return resp, nil
}
