package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListApPositionStatusRequest struct {
	JsonData string `position:"Query" name:"JsonData"`
}

func (r ListApPositionStatusRequest) Invoke(client *sdk.Client) (response *ListApPositionStatusResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ListApPositionStatusRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ListApPositionStatus", "", "")

	resp := struct {
		*responses.BaseResponse
		ListApPositionStatusResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ListApPositionStatusResponse

	err = client.DoAction(&req, &resp)
	return
}

type ListApPositionStatusResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
