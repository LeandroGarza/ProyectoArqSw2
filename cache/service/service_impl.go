package service

import (
	"cache/dto"
	"cache/service/repositories"
	"cache/utils/errors"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type CacheServiceImpl struct {
	Cacheclient repositories.CacheClient
}

func NewCacheServiceImpl(cacheclient repositories.CacheClient) *CacheServiceImpl {
	return &CacheServiceImpl{
		Cacheclient: cacheclient,
	}
}

func (s *CacheServiceImpl) GetUserData(id int) (dto.UserDto, errors.ApiError) {
	var userdto dto.UserDto

	userdata, err := s.Cacheclient.GetUserData(id)
	if err != nil {
		log.Debug("cache miss, searching in user service")

		resp, er := http.Get(fmt.Sprintf("http://localhost:9000/user/%v", id))
		if er != nil {
			return dto.UserDto{}, errors.NewNotFoundApiError("user not found in user service")
		}

		bytes, er := ioutil.ReadAll(resp.Body)
		if er != nil {
			return dto.UserDto{}, errors.NewInternalServerApiError("error reading body response", er)
		}
		resp.Body.Close()
		er = json.Unmarshal(bytes, &userdto)
		if er != nil {
			return dto.UserDto{}, errors.NewInternalServerApiError("error unmarshalling bytes", er)
		}

		userdata.Username = userdto.Username
		userdata.Email = userdto.Email
		_, er = s.Cacheclient.InsertUserData(id, userdata)
		if er != nil {
			return dto.UserDto{}, errors.NewInternalServerApiError("error inserting user in cache", er)
		}

		log.Debug("cache hit")
		return userdto, nil
	}

	userdto.Id = id
	userdto.Username = userdata.Username
	userdto.Email = userdata.Email

	return userdto, nil
}
