package request

import (
	"net/http"
	"sync"
)

var httpClientPool *sync.Pool
var httpClientOnce sync.Once

type ReturnHttpClient func(client *http.Client)
func init() {
	httpClientOnce.Do(func() {
		httpClientPool = &sync.Pool{
			New: func() interface{} {
				return &http.Client{
					Transport:     nil,
					CheckRedirect: nil,
					Jar:           nil,
					Timeout:       0,
				}
			},
		}
	})
}

func getClientFromPool() (*http.Client,ReturnHttpClient) {
	cli := httpClientPool.Get().(*http.Client)
	return cli,putClientToPool
}

func putClientToPool(cli *http.Client) {
	cli.Timeout = 0
	cli.Transport = nil
	cli.CheckRedirect = nil
	cli.Jar = nil

	httpClientPool.Put(cli)
}
