package repositories

import (
	"context"
	"fmt"
	"time"

	dtos "items/dtos"
	e "items/utils/errors"

	"github.com/karlseguin/ccache/v2"
)

type RepositoryCCache struct {
	Client     *ccache.Cache
	DefaultTTL time.Duration
}

func NewCCache(maxSize int64, itemsToPrune uint32, defaultTTL time.Duration) *RepositoryCCache {
	client := ccache.New(ccache.Configure().MaxSize(maxSize).ItemsToPrune(itemsToPrune))
	fmt.Println("[CCache] Initialized")
	return &RepositoryCCache{
		Client:     client,
		DefaultTTL: defaultTTL,
	}
}

func (repo *RepositoryCCache) Get(ctx context.Context, id string) (dtos.ItemDto, e.ApiError) {
	item := repo.Client.Get(id)
	if item == nil {
		return dtos.ItemDto{}, e.NewNotFoundApiError(fmt.Sprintf("item %s not found", id))
	}
	if item.Expired() {
		return dtos.ItemDto{}, e.NewNotFoundApiError(fmt.Sprintf("item %s not found", id))
	}
	return item.Value().(dtos.ItemDto), nil
}

func (repo *RepositoryCCache) InsertItem(ctx context.Context, item dtos.ItemDto) (dtos.ItemDto, e.ApiError) {
	repo.Client.Set(item.Id, item, repo.DefaultTTL)
	return item, nil
}

func (repo *RepositoryCCache) InsertItems(ctx context.Context, items dtos.ItemsDto) (dtos.ItemsDto, e.ApiError) {
	for _, item := range items {
		repo.Client.Set(item.Id, item, repo.DefaultTTL)
	}
	return items, nil
}

func (repo *RepositoryCCache) Update(ctx context.Context, item dtos.ItemDto) (dtos.ItemDto, e.ApiError) {
	repo.Client.Set(item.Id, item, repo.DefaultTTL)
	return item, nil
}

func (repo *RepositoryCCache) Delete(ctx context.Context, id string) e.ApiError {
	repo.Client.Delete(id)
	return nil
}

func (repo *RepositoryCCache) DeleteByUserId(ctx context.Context, userid int) e.ApiError {
	repo.Client.Clear()
	return nil
}
