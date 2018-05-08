package aegis

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyStrategyRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	RiskSubTypeName string `position:"Query" name:"RiskSubTypeName"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	CycleStartTime  string `position:"Query" name:"CycleStartTime"`
	Name            string `position:"Query" name:"Name"`
	CycleDays       string `position:"Query" name:"CycleDays"`
	Id              string `position:"Query" name:"Id"`
}

func (req *ModifyStrategyRequest) Invoke(client *sdk.Client) (resp *ModifyStrategyResponse, err error) {
	req.InitWithApiInfo("aegis", "2016-11-11", "ModifyStrategy", "vipaegis", "")
	resp = &ModifyStrategyResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyStrategyResponse struct {
	responses.BaseResponse
	RequestId      common.String
	Success        bool
	TotalCount     common.Integer
	HttpStatusCode common.Integer
	Result         ModifyStrategyResult
}

type ModifyStrategyResult struct {
	StrategyId common.Integer
}
