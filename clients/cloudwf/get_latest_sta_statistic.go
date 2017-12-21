package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetLatestStaStatisticRequest struct {
	ApgroupId int64 `position:"Query" name:"ApgroupId"`
}

func (r GetLatestStaStatisticRequest) Invoke(client *sdk.Client) (response *GetLatestStaStatisticResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetLatestStaStatisticRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetLatestStaStatistic", "", "")

	resp := struct {
		*responses.BaseResponse
		GetLatestStaStatisticResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetLatestStaStatisticResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetLatestStaStatisticResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
