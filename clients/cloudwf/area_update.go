package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AreaUpdateRequest struct {
	Name string `position:"Query" name:"Name"`
	Dids string `position:"Query" name:"Dids"`
	Aid  int64  `position:"Query" name:"Aid"`
	Sid  int64  `position:"Query" name:"Sid"`
}

func (r AreaUpdateRequest) Invoke(client *sdk.Client) (response *AreaUpdateResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AreaUpdateRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "AreaUpdate", "", "")

	resp := struct {
		*responses.BaseResponse
		AreaUpdateResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AreaUpdateResponse

	err = client.DoAction(&req, &resp)
	return
}

type AreaUpdateResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
