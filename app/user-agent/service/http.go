package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

type httpClient struct {
	client *http.Client
}

func newHttpClient() *httpClient {
	httpTransport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second, // 连接超时时间
			KeepAlive: 60 * time.Second, // 保持长连接的时间
		}).DialContext, // 设置连接的参数
		MaxIdleConns:          500,              // 最大空闲连接
		IdleConnTimeout:       60 * time.Second, // 空闲连接的超时时间
		ExpectContinueTimeout: 30 * time.Second, // 等待服务第一个响应的超时时间
		MaxIdleConnsPerHost:   100,              // 每个host保持的空闲连接数
	}

	return &httpClient{
		client: &http.Client{
			Transport: httpTransport,
		}}
}

func (hc *httpClient) Get(uri string, params map[string]string) (*http.Response, error) {
	return hc.requestWithoutBody(uri, params, nil, http.MethodGet)
}

func (hc *httpClient) Post(url string, body interface{}, params map[string]string) (*http.Response, error) {
	return hc.requestWithBody(url, body, params, nil, http.MethodPost)
}

func (hc *httpClient) Put(url string, body interface{}, params map[string]string) (*http.Response, error) {
	return hc.requestWithBody(url, body, params, nil, http.MethodPut)
}

func (hc *httpClient) Delete(uri string, params map[string]string) (*http.Response, error) {
	return hc.requestWithoutBody(uri, params, nil, http.MethodDelete)
}

func (hc *httpClient) requestWithoutBody(uri string, params map[string]string, headers map[string]string, method string) (*http.Response, error) {
	req, err := http.NewRequest(method, uri, nil)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	q := req.URL.Query()

	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}

	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}

	return hc.client.Do(req)
}

func (hc *httpClient) requestWithBody(url string, body interface{}, params map[string]string, headers map[string]string, method string) (*http.Response, error) {
	var bodyJSON []byte
	var req *http.Request
	if body != nil {
		var err error
		bodyJSON, err = json.Marshal(body)
		if err != nil {
			log.Println(err)
			return nil, errors.New("http post body to json failed")
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(bodyJSON))

	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf(fmt.Sprintf("new request is fail: %v \n", err))
	}

	if method == http.MethodPatch {
		req.Header.Set("Content-type", "application/json-patch+json")
	} else {
		req.Header.Set("Content-type", "application/json;charset=utf-8")
	}

	q := req.URL.Query()

	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}

	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}
	return hc.client.Do(req)
}
