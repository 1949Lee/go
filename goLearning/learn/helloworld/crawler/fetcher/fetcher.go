package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Fetch(url string) (body []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error get url: %s   ,status code: %d", url, resp.StatusCode)
	}

	all, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	return all, nil
}
