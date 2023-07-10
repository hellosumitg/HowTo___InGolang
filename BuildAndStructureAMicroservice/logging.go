package main

import (
	"context"
	"fmt"
	"time"
)

// LoggingService is a service wrapper that logs the execution time and errors of the underlying service.
type LoggingService struct {
	next Service
}

// NewLoggingService creates a new instance of LoggingService with the provided underlying Service.
func NewLoggingService(next Service) Service {
	return &LoggingService{
		next: next,
	}
}

// GetCatFact retrieves a cat fact and logs the execution time and any errors.
func (s *LoggingService) GetCatFact(ctx context.Context) (fact *CatFact, err error) {
	defer func(start time.Time) {
		fmt.Printf("fact=%v err=%s took=%v\n", fact.Fact, err, time.Since(start))
	}(time.Now())

	return s.next.GetCatFact(ctx)
}
