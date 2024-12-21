package post

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/rishu/microservice/config"
	"github.com/rishu/microservice/external/contants"
	"github.com/rishu/microservice/external/ohttp"
)

type Client interface {
	FetchPost(ctx context.Context, req *FetchPostRequest) (*FetchPostResponse, error)
}

type ClientImpl struct {
	httpRequestHandler ohttp.IHttpRequestHandler
	conf               *config.Config
}

var _ Client = &ClientImpl{}

func NewPostClientImpl(httpRequestHandler *ohttp.HttpRequestHandler, conf *config.Config) *ClientImpl {
	return &ClientImpl{
		httpRequestHandler: httpRequestHandler,
		conf:               conf,
	}
}

func (c *ClientImpl) FetchPost(ctx context.Context, req *FetchPostRequest) (*FetchPostResponse, error) {
	resp, err := c.httpRequestHandler.MakeHttpRequest(ctx, &ohttp.HttpRequest{
		Url:         fmt.Sprintf("%v/%v", c.conf.ExternalService.JsonPlaceholder.FetchPostUrl, req.PostId),
		Method:      contants.GetMethod,
		ContentType: contants.JsonContentType,
	})
	if err != nil {
		return nil, err
	}

	fetchPostResponse := &FetchPostResponse{}
	err = json.Unmarshal(resp.Body, fetchPostResponse)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	// Return the parsed response
	return fetchPostResponse, nil
}
