package post

import (
	"fmt"
	"github.com/rishu/microservice/external/contants"
	"github.com/rishu/microservice/external/pkg"
	placeholder "github.com/rishu/microservice/external/post/json_placeholder"
)

func (c *ClientImpl) NewPlaceholderRequest(req interface{}) pkg.SyncRequest {
	switch v := req.(type) {
	case *placeholder.FetchPostRequest:
		fetchReq := req.(*placeholder.FetchPostRequest)
		return &placeholder.VFetchPostRequest{
			Method: contants.GetMethod,
			Req:    fetchReq,
			Url:    c.conf.ExternalService.JsonPlaceholder.FetchPostUrl,
		}
	default:
		fmt.Println(v)
		return nil
	}
}
