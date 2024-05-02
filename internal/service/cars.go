package service

import (
	"effective_mobile_test_task/internal/repository/em"
	"effective_mobile_test_task/internal/repository/sql"
	"errors"
)

type CarsService struct {
	sql    *sql.CarsSQL
	emHttp *em.EmHttp
}

func newCarsService(sql *sql.CarsSQL, emHttp *em.EmHttp) *CarsService {
	return &CarsService{sql: sql, emHttp: emHttp}
}

func (s *CarsService) Create(regNums []string) error {
	for _, regNum := range regNums {
		car, err := s.emHttp.GetCarInfo(regNum)
		if err != nil {
			return err
		}
		_, err = s.sql.Create(car)
		if err != nil {
			logger.Error(err.Error())
			err = createPgError(err)
			return err
		}
	}
	return nil
}

func (s *CarsService) Edit(id int, editForm *sql.EditCar) error {
	if editForm == nil {
		return errors.New("Необходимо заполнить хотя бы один параметр в форме!")
	}
	_, err := s.sql.Edit(id, editForm)
	if err != nil {
		logger.Error(err.Error())
		err = editPgError(err, id)
		return err
	}
	return nil
}

func (s *CarsService) Delete(id int) error {
	_, err := s.sql.Delete(id)
	if err != nil {
		logger.Error(err.Error())
		err = deletePgError(err, id)
		return err
	}
	return nil
}

func (s *CarsService) GetAll(limit int, page int, selectForm *sql.CarFilter) (int, []sql.Car, error) {
	offset := (page - 1) * limit
	total, cars, err := s.sql.GetAll(limit, offset, selectForm)
	if err != nil {
		logger.Error(err.Error())
		return 0, nil, err
	}
	return total, cars, nil
}
