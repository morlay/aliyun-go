package aegis

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteStrategyRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	Id              string `position:"Query" name:"Id"`
}

func (req *DeleteStrategyRequest) Invoke(client *sdk.Client) (resp *DeleteStrategyResponse, err error) {
	req.InitWithApiInfo("aegis", "2016-11-11", "DeleteStrategy", "vipaegis", "")
	resp = &DeleteStrategyResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteStrategyResponse struct {
	responses.BaseResponse
	RequestId common.String
}
