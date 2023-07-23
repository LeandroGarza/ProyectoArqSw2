package services

import (
	"search/dtos"
	"search/utils/errors"
)

type Service interface {
	Search(query string) (dtos.ItemsSolrDto, error)
	SearchByUserId(id int) (dtos.ItemsSolrDto, error)
	InsertItems(itemsdto dtos.ItemsDto) (dtos.ItemsDto, errors.ApiError)
	DeleteAll() error
	DeleteByUserId(userid int) errors.ApiError
}
