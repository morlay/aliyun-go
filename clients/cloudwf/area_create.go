package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AreaCreateRequest struct {
	Name string `position:"Query" name:"Name"`
	Dids string `position:"Query" name:"Dids"`
	Sid  int64  `position:"Query" name:"Sid"`
}

func (r AreaCreateRequest) Invoke(client *sdk.Client) (response *AreaCreateResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AreaCreateRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "AreaCreate", "", "")

	resp := struct {
		*responses.BaseResponse
		AreaCreateResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AreaCreateResponse

	err = client.DoAction(&req, &resp)
	return
}

type AreaCreateResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
