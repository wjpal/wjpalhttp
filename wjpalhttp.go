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

	Post() ([]byte, error)
	Get() ([]byte, error)
}

func WJAPLHttp_NewHttpClient() WJAPLIHttpClient {
	return ihttp.NewHttpClient()
}
