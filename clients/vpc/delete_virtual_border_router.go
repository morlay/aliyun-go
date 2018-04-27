package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteVirtualBorderRouterRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	VbrId                string `position:"Query" name:"VbrId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteVirtualBorderRouterRequest) Invoke(client *sdk.Client) (resp *DeleteVirtualBorderRouterResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DeleteVirtualBorderRouter", "vpc", "")
	resp = &DeleteVirtualBorderRouterResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteVirtualBorderRouterResponse struct {
	responses.BaseResponse
	RequestId string
}
