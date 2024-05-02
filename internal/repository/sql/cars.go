package sql

import (
	"fmt"
	"strings"
)

const carsTable = "cars"

type People struct {
	Name       string  `json:"name"`
	Surname    string  `json:"surname"`
	Patronymic *string `json:"patronymic"`
}

type Car struct {
	Id     int    `json:"id" db:"id"`
	RegNum string `json:"RegNum" db:"reg_num"`
	Mark   string `json:"mark" db:"mark"`
	Model  string `json:"model" db:"model"`
	Year   *int   `json:"year" db:"year"`
	People People `json:"owner"`
}

type CarPreview struct {
	RegNum          string  `json:"reg_num"          db:"reg_num"`
	Mark            string  `json:"mark"             db:"mark"`
	Model           string  `json:"model"            db:"model"`
	Year            *int    `json:"year"             db:"year"`
	OwnerName       string  `json:"owner_name"       db:"owner_name"`
	OwnerSurname    string  `json:"owner_surname"    db:"owner_surname"`
	OwnerPatronymic *string `json:"owner_patronymic" db:"owner_patronymic"`
}

type CarFilter struct {
	RegNum          *string `form:"reg_num"`
	Mark            *string `form:"mark"`
	Model           *string `form:"model"`
	Year            *int    `form:"year"`
	OwnerName       *string `form:"owner_name"`
	OwnerSurname    *string `form:"owner_surname"`
	OwnerPatronymic *string `form:"owner_patronymic"`
}

type CreateCar struct {
	RegNum          string  `json:"reg_num"          db:"reg_num"`
	Mark            string  `json:"mark"             db:"mark"`
	Model           string  `json:"model"            db:"model"`
	Year            *int    `json:"year"             db:"year"`
	OwnerName       string  `json:"owner_name"       db:"owner_name"`
	OwnerSurname    string  `json:"owner_surname"    db:"owner_surname"`
	OwnerPatronymic *string `json:"owner_patronymic" db:"owner_patronymic"`
}

type EditCar struct {
	RegNum          *string `json:"reg_num"          db:"reg_num"`
	Mark            *string `json:"mark"             db:"mark"`
	Model           *string `json:"model"            db:"model"`
	Year            *int    `json:"year"             db:"year"`
	OwnerName       *string `json:"owner_name"       db:"owner_name"`
	OwnerSurname    *string `json:"owner_surname"    db:"owner_surname"`
	OwnerPatronymic *string `json:"owner_patronymic" db:"owner_patronymic"`
}

type CarsSQL struct {
	iBaseSQL[CarPreview]
}

func NewCarsSQL(dbPool *DbPool) *CarsSQL {
	sql := newIBaseSQL[CarPreview](dbPool, carsTable)
	return &CarsSQL{iBaseSQL: sql}
}

func (c *CarsSQL) Create(car *Car) (*CarPreview, error) {
	createForm := &CreateCar{
		RegNum:          car.RegNum,
		Mark:            car.Mark,
		Model:           car.Model,
		Year:            car.Year,
		OwnerName:       car.People.Name,
		OwnerSurname:    car.People.Surname,
		OwnerPatronymic: car.People.Patronymic,
	}
	return c.insert(createForm)
}

func (c *CarsSQL) Delete(id int) (*CarPreview, error) {
	return c.delete(id)
}

func (c *CarsSQL) Edit(id int, editForm *EditCar) (*CarPreview, error) {
	return c.update(id, editForm)
}

func (c *CarsSQL) GetAll(limit, offset int, selectForm *CarFilter) (int, []Car, error) {
	whereArray, args := c.createWhereStatement(selectForm)

	total, err := c.total(whereArray, args)
	if err != nil {
		return 0, nil, err
	}
	if total == 0 {
		return 0, nil, nil
	}

	cars, err := c.getAll(limit, offset, whereArray, args)
	if err != nil {
		return 0, nil, err
	}

	return total, cars, nil
}

func (c *CarsSQL) GetByID(id int) (*CarPreview, error) {
	return c.selectOne(id)
}

func (c *CarsSQL) createWhereStatement(selectForm *CarFilter) ([]string, []any) {
	if selectForm == nil {
		return nil, nil
	}
	var whereArray = make([]string, 0, 7)
	placeHolder := 1
	var args = make([]any, 0, 7)

	if selectForm.RegNum != nil {
		whereArray = append(whereArray, fmt.Sprintf("reg_num=$%d", placeHolder))
		args = append(args, *selectForm.RegNum)
		placeHolder++
	}
	if selectForm.Mark != nil {
		whereArray = append(whereArray, fmt.Sprintf("mark=$%d", placeHolder))
		args = append(args, *selectForm.Mark)
		placeHolder++
	}
	if selectForm.Model != nil {
		whereArray = append(whereArray, fmt.Sprintf("model=$%d", placeHolder))
		args = append(args, *selectForm.Model)
		placeHolder++
	}
	if selectForm.Year != nil {
		whereArray = append(whereArray, fmt.Sprintf("year=$%d", placeHolder))
		args = append(args, *selectForm.Year)
		placeHolder++
	}
	if selectForm.OwnerName != nil {
		whereArray = append(whereArray, fmt.Sprintf("owner_name=$%d", placeHolder))
		args = append(args, *selectForm.OwnerName)
		placeHolder++
	}
	if selectForm.OwnerSurname != nil {
		whereArray = append(whereArray, fmt.Sprintf("owner_surname=$%d", placeHolder))
		args = append(args, *selectForm.OwnerSurname)
		placeHolder++
	}
	if selectForm.OwnerPatronymic != nil {
		whereArray = append(whereArray, fmt.Sprintf("owner_patronymic=$%d", placeHolder))
		args = append(args, *selectForm.OwnerPatronymic)
		placeHolder++
	}
	return whereArray, args

}

func (c *CarsSQL) getAll(limit, offset int, whereArray []string, args []any) ([]Car, error) {
	var queryArray []string
	var cars []Car

	baseQuery := `
        SELECT
			id, 
			reg_num, 
			mark,
            model, 
			year, 
			owner_name,
            owner_surname, 
			owner_patronymic
        FROM cars
	`
	queryArray = append(queryArray, baseQuery)

	if whereArray != nil {
		where := "WHERE " + strings.Join(whereArray, " AND ")
		queryArray = append(queryArray, where)
	}
	if limit > 0 {
		limitQuery := fmt.Sprintf("LIMIT %d", limit)
		queryArray = append(queryArray, limitQuery)
	}
	if offset > 0 {
		offsetQuery := fmt.Sprintf("OFFSET %d", offset)
		queryArray = append(queryArray, offsetQuery)
	}

	query := strings.Join(queryArray, " ")

	rows, err := c.query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var car Car
		err = rows.Scan(
			&car.Id, &car.RegNum, &car.Mark, &car.Model, &car.Year,
			&car.People.Name, &car.People.Surname, &car.People.Patronymic,
		)
		cars = append(cars, car)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return cars, nil
}

func (c *CarsSQL) total(whereArray []string, args []any) (int, error) {
	var queryArray []string
	countQuery := `
        SELECT
			COUNT(*)
        FROM cars
	`
	queryArray = append(queryArray, countQuery)

	if whereArray != nil {
		where := "WHERE " + strings.Join(whereArray, " AND ")
		queryArray = append(queryArray, where)
	}

	totalQuery := strings.Join(queryArray, " ")

	var total int
	return total, c.queryRow(totalQuery, args...).Scan(&total)
}
