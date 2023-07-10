package main

import (
	"context"
	"encoding/json"
	"net/http"
)

// Service is the interface that defines the methods for retrieving a cat fact.
type Service interface {
	GetCatFact(context.Context) (*CatFact, error)
}

// CatFactService is a concrete implementation of the Service interface.
type CatFactService struct {
	url string
}

// NewCatFactService creates a new instance of CatFactService with the provided URL.
func NewCatFactService(url string) Service {
	return &CatFactService{
		url: url,
	}
}

// GetCatFact retrieves a cat fact from the specified URL.
func (s *CatFactService) GetCatFact(ctx context.Context) (*CatFact, error) {
	res, err := http.Get(s.url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	fact := &CatFact{}
	if err := json.NewDecoder(res.Body).Decode(fact); err != nil {
		return nil, err
	}

	return fact, nil
}
