package isdayoff

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const (
	baseURL = "https://isdayoff.ru/"
)

func IsDayOff(httpCli *http.Client, date time.Time) (bool, error) {
	res, err := httpCli.Get(baseURL + date.Format("20060102"))
	if err != nil {
		return false, err
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return false, err
	}
	if res.StatusCode != http.StatusOK {
		return false, fmt.Errorf("Unexpected response: %s", res.Status)
	}
	return strconv.ParseBool(string(body))
}

func IsDayOffToday(httpCli *http.Client) (bool, error) {
	return IsDayOff(httpCli, time.Now())
}
