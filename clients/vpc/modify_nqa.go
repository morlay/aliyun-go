package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyNqaRequest struct {
	requests.RpcRequest
	DestinationIp        string `position:"Query" name:"DestinationIp"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	NqaId                string `position:"Query" name:"NqaId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifyNqaRequest) Invoke(client *sdk.Client) (resp *ModifyNqaResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "ModifyNqa", "vpc", "")
	resp = &ModifyNqaResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyNqaResponse struct {
	responses.BaseResponse
	RequestId common.String
}
