package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetApDetailedStatusRequest struct {
	Id int64 `position:"Query" name:"Id"`
}

func (r GetApDetailedStatusRequest) Invoke(client *sdk.Client) (response *GetApDetailedStatusResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetApDetailedStatusRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetApDetailedStatus", "", "")

	resp := struct {
		*responses.BaseResponse
		GetApDetailedStatusResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetApDetailedStatusResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetApDetailedStatusResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
