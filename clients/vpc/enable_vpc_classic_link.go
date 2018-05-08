package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type EnableVpcClassicLinkRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	VpcId                string `position:"Query" name:"VpcId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *EnableVpcClassicLinkRequest) Invoke(client *sdk.Client) (resp *EnableVpcClassicLinkResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "EnableVpcClassicLink", "vpc", "")
	resp = &EnableVpcClassicLinkResponse{}
	err = client.DoAction(req, resp)
	return
}

type EnableVpcClassicLinkResponse struct {
	responses.BaseResponse
	RequestId common.String
}
