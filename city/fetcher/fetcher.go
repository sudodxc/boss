package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//fetcher all city
func Fetch(url string) ([]byte, error) {
	// first get all page
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {

		return nil, fmt.Errorf("wrong stats code :%d", res.StatusCode)
	}
	all, _ := ioutil.ReadAll(res.Body)
	return all, nil

}
