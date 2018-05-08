package aegis

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyBatchIgnoreVulRequest struct {
	requests.RpcRequest
	Reason          string `position:"Query" name:"Reason"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	Info            string `position:"Query" name:"Info"`
}

func (req *ModifyBatchIgnoreVulRequest) Invoke(client *sdk.Client) (resp *ModifyBatchIgnoreVulResponse, err error) {
	req.InitWithApiInfo("aegis", "2016-11-11", "ModifyBatchIgnoreVul", "vipaegis", "")
	resp = &ModifyBatchIgnoreVulResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyBatchIgnoreVulResponse struct {
	responses.BaseResponse
	RequestId      common.String
	Success        bool
	HttpStatusCode common.Integer
}
