package ohttp

import (
	"bytes"
	"context"
	"fmt"
	"github.com/rishu/microservice/external/contants"
	"github.com/rishu/microservice/external/pkg"
	"io/ioutil"
	"net/http"
	"net/url"
)

type IHttpRequestHandler interface {
	MakeHttpRequest(context.Context, pkg.SyncRequest) (interface{}, error)
}

type HttpRequestHandler struct {
	client *http.Client
}

func NewHttpRequestHandler(client *http.Client) *HttpRequestHandler {
	return &HttpRequestHandler{
		client: client,
	}
}

func (h *HttpRequestHandler) MakeHttpRequest(ctx context.Context, request pkg.SyncRequest) (interface{}, error) {
	uri, err := url.Parse(request.GetURL())
	if err != nil {
		return nil, fmt.Errorf("URL could not be parsed: %w", err)
	}
	reqBody, err := request.Marshal()
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequest(request.GetMethod(), uri.String(), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("error encountered creating http httpRequest: %w", err)
	}

	// Set the appropriate headers
	httpReq.Header.Set("Content-Type", contants.JsonContentType)

	// Use http.Client to send the request
	response, err := h.client.Do(httpReq)
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

	resp, err := request.GetResponse().Unmarshal(data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
