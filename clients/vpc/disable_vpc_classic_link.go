package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DisableVpcClassicLinkRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	VpcId                string `position:"Query" name:"VpcId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DisableVpcClassicLinkRequest) Invoke(client *sdk.Client) (response *DisableVpcClassicLinkResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DisableVpcClassicLinkRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "DisableVpcClassicLink", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		DisableVpcClassicLinkResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DisableVpcClassicLinkResponse

	err = client.DoAction(&req, &resp)
	return
}

type DisableVpcClassicLinkResponse struct {
	RequestId string
}
