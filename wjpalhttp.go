package wjpalhttp

import ihttp "github.com/wjpal/wjpalhttp/internal"

func WJPALHttp_Initialize() error {
	return ihttp.IHttp_initialize()
}

type WJAPLIHttpClient interface {
	Clean()
	PutUrl(url string)
	PutBody(body []byte)
	AppendHeader(headKey, headValue string)
	GetLastStatusCode() int

	Post() ([]byte, error)
	Get() ([]byte, error)
	Put() ([]byte, error)
	Request(string) ([]byte, error) // request anymethod
}

func WJAPLHttp_NewHttpClient() WJAPLIHttpClient {
	return ihttp.NewHttpClient()
}
