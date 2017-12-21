package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetLatestApStatisticRequest struct {
	ApgroupId int64 `position:"Query" name:"ApgroupId"`
}

func (r GetLatestApStatisticRequest) Invoke(client *sdk.Client) (response *GetLatestApStatisticResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetLatestApStatisticRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetLatestApStatistic", "", "")

	resp := struct {
		*responses.BaseResponse
		GetLatestApStatisticResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetLatestApStatisticResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetLatestApStatisticResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
