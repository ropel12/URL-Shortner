package repository

import (
	"sync"

	"github.com/ropel12/URL-Shortner/entity"
)

type URLRepository struct {
	mu   sync.Mutex
	Data map[string]string
}

func NewMapRepository() URLRepository {
	data := make(map[string]string, 0)
	return URLRepository{
		Data: data,
	}
}

func (r *URLRepository) Get(path string) (*entity.URL, error) {
	data, ok := r.Data[path]
	if !ok {
		return nil, entity.ErrURLNotFound
	}
	return &entity.URL{
		LongURL:  data,
		ShortURL: path,
	}, nil
}

func (r *URLRepository) Create(longURL string) (*entity.URL, error) {
	r.mu.Lock()
	res, err := r.Get(longURL)
	if err == nil {
		return res, entity.ErrCustomURLIsExists
	}
	shorturl := entity.GetRandomShortURL(longURL)
	r.Data[shorturl] = longURL
	defer r.mu.Unlock()
	return &entity.URL{
			LongURL:  longURL,
			ShortURL: shorturl,
		},
		nil

}

func (r *URLRepository) CreateCustom(longURL, customPath string) (*entity.URL, error) {
	r.mu.Lock()
	res, err := r.Get(customPath)
	if err == nil {
		return res, entity.ErrCustomURLIsExists
	}
	r.Data[customPath] = longURL
	defer r.mu.Unlock()
	return &entity.URL{
			LongURL:  longURL,
			ShortURL: customPath,
		},
		nil
}
