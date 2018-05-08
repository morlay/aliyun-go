package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DisableVpcClassicLinkRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	VpcId                string `position:"Query" name:"VpcId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DisableVpcClassicLinkRequest) Invoke(client *sdk.Client) (resp *DisableVpcClassicLinkResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DisableVpcClassicLink", "vpc", "")
	resp = &DisableVpcClassicLinkResponse{}
	err = client.DoAction(req, resp)
	return
}

type DisableVpcClassicLinkResponse struct {
	responses.BaseResponse
	RequestId common.String
}
