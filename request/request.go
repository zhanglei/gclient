package request

import (
	"encoding"
	"errors"
	"github.com/jianzhiyao/gclient/consts"
	"github.com/jianzhiyao/gclient/consts/content_type"
	"net/http"
)

type Request struct {
	method  string
	url     string
	headers map[string]string
	body    []byte
}

func New(method, url string) (*Request, error) {
	switch method {
	case http.MethodGet:
	case http.MethodPost:
	case http.MethodConnect:
	case http.MethodDelete:
	case http.MethodHead:
	case http.MethodOptions:
	case http.MethodPatch:
	case http.MethodPut:
	case http.MethodTrace:
	default:
		return nil, errors.New("not a valid http method")
	}
	return &Request{
		method:  method,
		url:     url,
		headers: make(map[string]string),
	}, nil
}

func (r *Request) SetHeader(key string, value string) {
	r.headers[key] = value
}

func (r *Request) Method() string {
	return r.method
}

func (r *Request) Headers() map[string]string {
	return r.headers
}

func (r *Request) Json(body interface{}) (err error) {
	if e := r.Body(body); e != nil {
		return e
	}
	r.SetHeader(consts.HeaderContentType, content_type.ApplicationJson)
	return
}

func (r *Request) Xml(body interface{}) (err error) {
	if e := r.Body(body); e != nil {
		return e
	}
	r.SetHeader(consts.HeaderContentType, content_type.ApplicationXml)
	return
}

func (r *Request) MultiForm(body interface{}) (err error) {
	if e := r.Body(body); e != nil {
		return e
	}
	r.SetHeader(consts.HeaderContentType, content_type.MultipartFormData)
	return
}

func (r *Request) Form(body interface{}) (err error) {
	if e := r.Body(body); e != nil {
		return e
	}
	r.SetHeader(consts.HeaderContentType, content_type.ApplicationXWwwFormUrlencoded)
	return
}

func (r *Request) Body(body interface{}) (err error) {
	switch body := body.(type) {
	case []byte:
		r.body = body
	case string:
		r.body = []byte(body)
	case encoding.BinaryMarshaler:
		r.body, err = body.MarshalBinary()
	default:
		err = ErrCanNotMarshal
	}

	return
}
