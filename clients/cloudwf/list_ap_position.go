package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListApPositionRequest struct {
	MapId int64 `position:"Query" name:"MapId"`
}

func (r ListApPositionRequest) Invoke(client *sdk.Client) (response *ListApPositionResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ListApPositionRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ListApPosition", "", "")

	resp := struct {
		*responses.BaseResponse
		ListApPositionResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ListApPositionResponse

	err = client.DoAction(&req, &resp)
	return
}

type ListApPositionResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
