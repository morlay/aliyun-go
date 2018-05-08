package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateNqaRequest struct {
	requests.RpcRequest
	DestinationIp        string `position:"Query" name:"DestinationIp"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	RouterId             string `position:"Query" name:"RouterId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *CreateNqaRequest) Invoke(client *sdk.Client) (resp *CreateNqaResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "CreateNqa", "vpc", "")
	resp = &CreateNqaResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateNqaResponse struct {
	responses.BaseResponse
	RequestId common.String
	NqaId     common.String
}
