package repositories

import "net/http"

type Client interface {
	Search(query string) (*http.Response, error)
	SearchByUserId(id int) (*http.Response, error)
	Delete(userid int) (*http.Response, error)
	DeleteAll() (*http.Response, error)
}
