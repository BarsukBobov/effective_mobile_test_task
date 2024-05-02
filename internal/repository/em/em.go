package em

import (
	"effective_mobile_test_task/internal/repository/sql"
	"effective_mobile_test_task/pkg/misc"
	"encoding/json"
	"net/http"
)

var logger = misc.GetLogger()

type Links struct {
	CarInfo string
}

type EmHttp struct {
	*misc.HttpAdapter
	links Links
}

func NewEmHttp(baseUrl string) (*EmHttp, error) {
	httpAdapter := misc.NewHttpAdapter(baseUrl)

	links := Links{
		CarInfo: baseUrl,
	}
	return &EmHttp{
		HttpAdapter: httpAdapter,
		links:       links,
	}, nil
}

func (e *EmHttp) GetCarInfo(regNum string) (*sql.Car, error) {
	req, err := http.NewRequest("GET", e.links.CarInfo, nil)
	if err != nil {
		return nil, err
	}

	query := req.URL.Query()
	query.Add("regNum", regNum)
	req.URL.RawQuery = query.Encode()

	resp, err := e.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var car sql.Car
	if err := json.NewDecoder(resp.Body).Decode(&car); err != nil {
		return nil, err
	}
	return &car, nil
}
