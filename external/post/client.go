package post

import (
	"context"
	"fmt"
	"github.com/rishu/microservice/config"
	"github.com/rishu/microservice/external/enums"
	"github.com/rishu/microservice/external/ohttp"
	"github.com/rishu/microservice/external/pkg"
	"github.com/rishu/microservice/external/post/json_placeholder"
)

type Client interface {
	FetchPost(ctx context.Context, req *placeholder.FetchPostRequest) (*placeholder.FetchPostResponse, error)
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

func (c *ClientImpl) requestFactoryMap() map[enums.Vendor]pkg.SyncRequestFactory {
	return map[enums.Vendor]pkg.SyncRequestFactory{
		enums.JsonPlaceholder: c.NewPlaceholderRequest,
	}
}

func (c *ClientImpl) FetchPost(ctx context.Context, req *placeholder.FetchPostRequest) (*placeholder.FetchPostResponse, error) {
	venReq, err := pkg.NewVendorRequest(req, c.requestFactoryMap())
	if err != nil {
		return nil, err
	}
	fmt.Println(venReq)
	resp, err := c.httpRequestHandler.MakeHttpRequest(ctx, venReq)
	if err != nil {
		return nil, err
	}

	return resp.(*placeholder.FetchPostResponse), nil
}
