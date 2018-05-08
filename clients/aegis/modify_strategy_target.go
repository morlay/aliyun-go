package aegis

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyStrategyTargetRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	Type            string `position:"Query" name:"Type"`
	Config          string `position:"Query" name:"Config"`
	Target          string `position:"Query" name:"Target"`
}

func (req *ModifyStrategyTargetRequest) Invoke(client *sdk.Client) (resp *ModifyStrategyTargetResponse, err error) {
	req.InitWithApiInfo("aegis", "2016-11-11", "ModifyStrategyTarget", "vipaegis", "")
	resp = &ModifyStrategyTargetResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyStrategyTargetResponse struct {
	responses.BaseResponse
	RequestId common.String
}
