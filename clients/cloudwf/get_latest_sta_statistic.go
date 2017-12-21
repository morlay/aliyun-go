package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetLatestStaStatisticRequest struct {
	requests.RpcRequest
	ApgroupId int64 `position:"Query" name:"ApgroupId"`
}

func (req *GetLatestStaStatisticRequest) Invoke(client *sdk.Client) (resp *GetLatestStaStatisticResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetLatestStaStatistic", "", "")
	resp = &GetLatestStaStatisticResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetLatestStaStatisticResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
