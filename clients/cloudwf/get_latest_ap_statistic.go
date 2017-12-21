package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetLatestApStatisticRequest struct {
	requests.RpcRequest
	ApgroupId int64 `position:"Query" name:"ApgroupId"`
}

func (req *GetLatestApStatisticRequest) Invoke(client *sdk.Client) (resp *GetLatestApStatisticResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetLatestApStatistic", "", "")
	resp = &GetLatestApStatisticResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetLatestApStatisticResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
