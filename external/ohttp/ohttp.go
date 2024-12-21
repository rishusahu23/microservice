package ohttp

import (
	"context"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"strings"
)

type IHttpRequestHandler interface {
	MakeHttpRequest(context.Context, *HttpRequest) (*HttpResponse, error)
}

type HttpRequestHandler struct {
	client *http.Client
}

type HttpRequest struct {
	Body        []byte
	Url         string
	Method      string
	ContentType string
}

type HttpResponse struct {
	Body       []byte
	StatusCode int
}

func getHttpClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
}

func NewHttpRequestHandler(client *http.Client) *HttpRequestHandler {
	return &HttpRequestHandler{
		client: client,
	}
}

func (h *HttpRequestHandler) MakeHttpRequest(ctx context.Context, request *HttpRequest) (*HttpResponse, error) {
	req, err := http.NewRequest(request.Method, request.Url, strings.NewReader(string(request.Body)))
	if err != nil {
		//logger.Error(ctx, "The HTTP request creation failed with error: %s\n", zap.Error(err))
		return nil, err
	}

	// Set the appropriate headers
	req.Header.Set("Content-Type", request.ContentType)

	// Use http.Client to send the request
	response, err := h.client.Do(req)
	if err != nil {
		//logger.Error(ctx, "The HTTP request failed with error: %s\n", zap.Error(err))
		return nil, err
	}
	defer func() {
		if err = response.Body.Close(); err != nil {
			//logger.Error(ctx, "error closing response body: %v", zap.Error(err))
		}
	}()
	// Read the response body
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		//logger.Error(ctx, "Failed to read the response body: %s\n", zap.Error(err))
		return nil, err
	}

	return &HttpResponse{
		Body: data,
	}, nil
}
