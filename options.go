package gclient

import (
	"context"
	"github.com/jianzhiyao/gclient/consts"
	"net/http"
	"time"
)

type Option func(req *Client)
type CheckRedirectHandler func(req *http.Request, via []*http.Request) error

func OptTimeout(timeout time.Duration) Option {
	return func(req *Client) {
		req.clientTimeout = timeout
	}
}

func OptContext(ctx context.Context) Option {
	return func(req *Client) {
		req.ctx = ctx
	}
}

func OptHeader(key string, value ...string) Option {
	return func(req *Client) {
		req.headers[key] = value
	}
}

func OptUserAgent(ua string) Option {
	return OptHeader(consts.HeaderUserAgent, ua)
}

func OptHeaders(headers map[string][]string) Option {
	return func(req *Client) {
		for key, value := range headers {
			req.headers[key] = value
		}
	}
}

func OptEnableGzip() Option {
	return enableSign(SignGzip)
}

func OptDisableGzip() Option {
	return disableSign(SignGzip)
}

func OptEnableBr() Option {
	return enableSign(SignBr)
}

func OptDisableBr() Option {
	return disableSign(SignBr)
}

func OptCookieJar(jar http.CookieJar) Option {
	return func(req *Client) {
		req.clientCookieJar = jar
	}
}

func OptTransport(roundTripper http.RoundTripper) Option {
	return func(req *Client) {
		req.clientTransport = roundTripper
	}
}

func OptCheckRedirectHandler(clientCheckRedirect CheckRedirectHandler) Option {
	return func(req *Client) {
		req.clientCheckRedirect = clientCheckRedirect
	}
}

//OptRetry set retry num of requests in one client
func OptRetry(num int) Option {
	return func(req *Client) {
		req.retry = num
	}
}

func enableSign(t Sign) Option {
	return func(req *Client) {
		req.sign |= int8(t)

		var contentEncoding []string
		if req.sign&int8(SignGzip) != 0 {
			contentEncoding = append(contentEncoding, consts.ContentEncodingGzip)
		}
		if req.sign&int8(SignBr) != 0 {
			contentEncoding = append(contentEncoding, consts.ContentEncodingBr)
		}
		if len(contentEncoding) > 0 {
			req.headers[consts.HeaderAcceptEncoding] = contentEncoding
		}
	}
}

func disableSign(t Sign) Option {
	return func(req *Client) {
		req.sign ^= int8(t)

		var contentEncoding []string
		if req.sign&int8(SignGzip) != 0 {
			contentEncoding = append(contentEncoding, consts.ContentEncodingGzip)
		}
		if req.sign&int8(SignBr) != 0 {
			contentEncoding = append(contentEncoding, consts.ContentEncodingBr)
		}
		if len(contentEncoding) > 0 {
			req.headers[consts.HeaderAcceptEncoding] = contentEncoding
		}
	}
}
