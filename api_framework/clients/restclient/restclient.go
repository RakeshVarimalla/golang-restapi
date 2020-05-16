package restclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

var (
	enabledMocks = false
	mocks = make(map[string]*Mock)
)
type Mock struct{
	Url string
	HttpMethod string
	Response *http.Response
	Err error
}
func FlushMocks(){
	mocks = make(map[string]*Mock)
}
func StartMocks(){
	enabledMocks = true
}

func StopMocks(){
	enabledMocks = false
}

func AddMocks(mock Mock){
	mocks[GetMockId(mock.HttpMethod, mock.Url)] = &mock
}

func GetMockId(httpMethod string, Url string) string{
return fmt.Sprintf("%s_%s", httpMethod,Url)
}
func Post(url string, body interface{}, headers http.Header)(*http.Response, error) {

	if enabledMocks {
		mock := mocks[GetMockId(http.MethodPost,url)]
		if mock == nil {
			fmt.Println("No mockup found for given request")
			return nil, errors.New("No mockup found for given request ")
		}else{
			return  mock.Response, mock.Err
		}
	}

	jsonValue, err := json.Marshal(body)
	log.Println(string(jsonValue))
	log.Println(url)
	log.Println(headers)
	if err != nil {
		return nil, err
	}
     request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonValue))
     request.Header = headers
     client :=  http.Client{}
     return client.Do(request)
}
