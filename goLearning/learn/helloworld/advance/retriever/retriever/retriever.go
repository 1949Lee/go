package retriever

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	UserAgent string
	TimeOut   time.Duration
	Name      string
}

func (r *Retriever) String() string {
	s := fmt.Sprintf("Retriever: {Name:%s}", r.Name)
	return s
}

func (r *Retriever) Post(name string) string {
	r.Name = name
	return r.Name
}

func (r *Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	result, err := httputil.DumpResponse(resp, true)

	resp.Body.Close()
	if err != nil {
		panic(err)
	}

	return string(result)

}
