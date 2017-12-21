package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteNqaRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	NqaId                string `position:"Query" name:"NqaId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteNqaRequest) Invoke(client *sdk.Client) (resp *DeleteNqaResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DeleteNqa", "vpc", "")
	resp = &DeleteNqaResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteNqaResponse struct {
	responses.BaseResponse
	RequestId string
}
