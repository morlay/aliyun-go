package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DelApPositionRequest struct {
	ApAssetId int64 `position:"Query" name:"ApAssetId"`
	MapId     int64 `position:"Query" name:"MapId"`
}

func (r DelApPositionRequest) Invoke(client *sdk.Client) (response *DelApPositionResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DelApPositionRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "DelApPosition", "", "")

	resp := struct {
		*responses.BaseResponse
		DelApPositionResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DelApPositionResponse

	err = client.DoAction(&req, &resp)
	return
}

type DelApPositionResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
