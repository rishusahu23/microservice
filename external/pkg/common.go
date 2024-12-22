package pkg

import (
	"fmt"
	"github.com/rishu/microservice/external/enums"
)

type SyncRequest interface {
	GetMethod() string
	GetURL() string
	GetResponse() Response
	Marshal() ([]byte, error)
}

type Response interface {
	Unmarshal(b []byte) (interface{}, error)
}

type RequestWithHeader interface {
	GetHeader() enums.Vendor
}

type SyncRequestFactory func(req interface{}) SyncRequest

func NewVendorRequest(req RequestWithHeader, reqFactMap map[enums.Vendor]SyncRequestFactory) (SyncRequest, error) {
	v := req.GetHeader()
	vendor, ok := reqFactMap[v]
	if !ok {
		return nil, fmt.Errorf("invalid error %v", v)
	}
	vendorReq := vendor(req)
	if vendorReq == nil {
		return nil, fmt.Errorf("invalid request %v", req)
	}
	return vendorReq, nil
}
