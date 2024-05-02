package service

import (
	"effective_mobile_test_task/internal/repository/em"
	"effective_mobile_test_task/internal/repository/sql"
	"effective_mobile_test_task/pkg/misc"
)

var logger = misc.GetLogger()

type Service struct {
	Cars *CarsService
}

func NewService(sql *sql.SQL, emHttp *em.EmHttp) *Service {
	return &Service{
		Cars: newCarsService(sql.Cars, emHttp),
	}
}
