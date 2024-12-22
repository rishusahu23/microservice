package placeholder

import "github.com/rishu/microservice/external/enums"

type FetchPostRequest struct {
	Vendor enums.Vendor
	PostId string
}

func (f *FetchPostRequest) GetHeader() enums.Vendor {
	return f.Vendor
}
