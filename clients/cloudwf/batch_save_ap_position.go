package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type BatchSaveApPositionRequest struct {
	JsonData string `position:"Query" name:"JsonData"`
}

func (r BatchSaveApPositionRequest) Invoke(client *sdk.Client) (response *BatchSaveApPositionResponse, err error) {
	req := struct {
		*requests.RpcRequest
		BatchSaveApPositionRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "BatchSaveApPosition", "", "")

	resp := struct {
		*responses.BaseResponse
		BatchSaveApPositionResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.BatchSaveApPositionResponse

	err = client.DoAction(&req, &resp)
	return
}

type BatchSaveApPositionResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
