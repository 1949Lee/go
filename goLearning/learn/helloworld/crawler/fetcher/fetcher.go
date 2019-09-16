package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var client = http.Client{}

var rateLimiter = time.Tick(10 * time.Millisecond)

func Fetch(url string) (body []byte, err error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	<-rateLimiter
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.142 Safari/537.36")
	resp, err := client.Do(request)
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
