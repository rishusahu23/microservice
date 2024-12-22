package placeholder

import (
	"encoding/json"
	"fmt"
	"github.com/rishu/microservice/external/contants"
	"github.com/rishu/microservice/external/pkg"
)

type VFetchPostRequest struct {
	Method string
	Req    *FetchPostRequest
	Url    string
}

func (v *VFetchPostRequest) GetMethod() string {
	return contants.GetMethod
}

func (v *VFetchPostRequest) GetURL() string {
	return fmt.Sprintf("%v/%v", v.Url, v.Req.PostId)
}

func (v *VFetchPostRequest) GetResponse() pkg.Response {
	return &VFetchPostResponse{}
}

func (v *VFetchPostRequest) Marshal() ([]byte, error) {
	return nil, nil
}

type VFetchPostResponse struct {
}

func (V *VFetchPostResponse) Unmarshal(b []byte) (interface{}, error) {
	vendirRes := &FetchPostResponse{}
	err := json.Unmarshal(b, vendirRes)
	if err != nil {
		return nil, err
	}
	return vendirRes, nil
}
