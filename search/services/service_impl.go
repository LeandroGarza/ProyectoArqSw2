package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	dtos "search/dtos"
	clients "search/services/repositories"
	"search/utils/errors"
)

type SearchService struct {
	searchclient *clients.SearchClient
}

func NewSearchService(searchclient *clients.SearchClient) *SearchService {
	return &SearchService{
		searchclient: searchclient,
	}
}

func (s *SearchService) Search(query string) (dtos.ItemsSolrDto, error) {
	// llamada al search de clients
	r, err := s.searchclient.Search(query)
	if err != nil {
		return dtos.ItemsSolrDto{}, err
	}

	// lectura de la response y guardamos los bytes
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return dtos.ItemsSolrDto{}, err
	}

	// llamada a la funcion que parsea items
	rdto, err := parseItems(bytes)
	if err != nil {
		return dtos.ItemsSolrDto{}, err
	}

	return rdto.Response.Docs, nil
}

func (s *SearchService) SearchByUserId(id int) (dtos.ItemsSolrDto, error) {
	r, err := s.searchclient.SearchByUserId(id)
	if err != nil {
		return dtos.ItemsSolrDto{}, err
	}

	// lectura de la response y guardamos los bytes
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return dtos.ItemsSolrDto{}, err
	}

	// llamada a la funcion que parsea items
	rdto, err := parseItems(bytes)
	if err != nil {
		return dtos.ItemsSolrDto{}, err
	}

	return rdto.Response.Docs, nil
}

func (s *SearchService) InsertItems(itemsdto dtos.ItemsDto) (dtos.ItemsDto, errors.ApiError) {
	url := "http://localhost:8983/solr/items/update/json/docs?commit=true"

	var itemssolrdto dtos.ItemsSolrDto
	for _, itemdto := range itemsdto {
		var itemsolr dtos.ItemSolrDto

		itemsolr.Id = itemdto.Id
		itemsolr.Title = append(itemsolr.Title, itemdto.Title)
		itemsolr.Userid = append(itemsolr.Userid, itemdto.Userid)
		itemsolr.Image = append(itemsolr.Image, itemdto.Image)
		itemsolr.Currency = append(itemsolr.Currency, itemdto.Currency)
		itemsolr.Price = append(itemsolr.Price, itemdto.Price)
		itemsolr.Sale_sate = append(itemsolr.Sale_sate, itemdto.Sale_sate)
		itemsolr.Condition = append(itemsolr.Condition, itemdto.Condition)
		itemsolr.Address = append(itemsolr.Address, itemdto.Address)

		itemssolrdto = append(itemssolrdto, itemsolr)
	}

	// Convertir el objeto itemssolrdto a JSON
	requestBody, err := json.Marshal(itemssolrdto)
	if err != nil {
		return dtos.ItemsDto{}, errors.NewInternalServerApiError("Error al convertir los items en JSON", err)
	}

	// Realizar la solicitud HTTP POST
	response, err := http.Post(url, "application/json", bytes.NewReader(requestBody))
	if err != nil {
		return dtos.ItemsDto{}, errors.NewInternalServerApiError("Error al realizar la solicitud HTTP a Solr", err)
	}
	defer response.Body.Close()

	// Verificar el código de respuesta
	if response.StatusCode != http.StatusOK {
		// Leer el cuerpo de la respuesta
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return dtos.ItemsDto{}, errors.NewInternalServerApiError("Error al leer la respuesta de la solicitud HTTP", err)
		}
		return dtos.ItemsDto{}, errors.NewInternalServerApiError(fmt.Sprintf("Error al insertar los items en Solr. Código de respuesta: %d. Mensaje: %s", response.StatusCode, string(body)), nil)
	}

	return itemsdto, nil
}

func (s *SearchService) DeleteAll() error {
	_, err := s.searchclient.DeleteAll()
	if err != nil {
		return errors.NewInternalServerApiError("error deleting documents", err)
	}

	return nil
}

func (s *SearchService) DeleteByUserId(userid int) errors.ApiError {
	_, err := s.searchclient.DeleteByUserId(userid)
	if err != nil {
		return errors.NewInternalServerApiError(fmt.Sprintf("Error deleting documents by userid: %v", userid), err)
	}

	return nil
}

func parseItems(bytes []byte) (dtos.ResponseDto, error) {
	var rdto dtos.ResponseDto
	// decodificamos los bytes y los guardamos en la variable que
	if err := json.Unmarshal(bytes, &rdto); err != nil {
		return dtos.ResponseDto{}, err
	}
	// Imprimir el contenido de rdto para verificarlo
	fmt.Printf("Contenido de rdto: %+v\n", rdto)
	fmt.Printf("Contenido de Docs de rdto: %+v\n", rdto.Response.Docs)

	return rdto, nil
}
