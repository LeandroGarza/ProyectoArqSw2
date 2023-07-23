package repositories

import (
	"encoding/json"
	"fmt"

	"cache/model"

	e "cache/utils/errors"

	"github.com/bradfitz/gomemcache/memcache"
	log "github.com/sirupsen/logrus"
)

type CacheClientImpl struct {
	cacheclient *memcache.Client
}

func NewCacheClientImpl(cachehost string, cacheport int) *CacheClientImpl {
	cacheclient := memcache.New(fmt.Sprintf("%v:%v", cachehost, cacheport))
	if cacheclient == nil {
		log.Panic("error creating new cache instance")
	}
	return &CacheClientImpl{
		cacheclient: cacheclient,
	}
}

func (c *CacheClientImpl) InsertUserData(id int, userdata model.Data) (model.Data, error) {
	bytesdata, err := json.Marshal(userdata)
	if err != nil {
		return model.Data{}, e.NewInternalServerApiError("Error in data marshall", err)
	}

	// agregar nuevo item con una expiracion de 1200 segundos
	er := c.cacheclient.Set(&memcache.Item{
		Key:        fmt.Sprintf("%v", id),
		Value:      bytesdata,
		Expiration: 1200,
	})
	if er != nil {
		return model.Data{}, e.NewInternalServerApiError("Error inserting userdata to cache", err)
	}
	return userdata, nil
}

func (c *CacheClientImpl) GetUserData(id int) (model.Data, error) {
	item, err := c.cacheclient.Get(fmt.Sprintf("%v", id))
	if err != nil {
		return model.Data{}, e.NewBadRequestApiError("cache miss")
	}
	var data model.Data
	er := json.Unmarshal(item.Value, &data)
	if er != nil {
		return model.Data{}, e.NewInternalServerApiError("error unmarshalling data", err)
	}

	return data, nil
}

func (c *CacheClientImpl) DeleteUserData(id int) error {
	err := c.cacheclient.Delete(fmt.Sprintf("%v", id))
	if err != nil {
		return e.NewInternalServerApiError("error deleting user data", err)
	}
	return nil
}
