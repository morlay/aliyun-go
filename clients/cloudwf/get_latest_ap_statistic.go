package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId common.String
	Success   bool
	Message   common.String
	Data      common.String
	ErrorCode common.Integer
	ErrorMsg  common.String
}
