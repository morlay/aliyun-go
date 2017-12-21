package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type EnableVpcClassicLinkRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	VpcId                string `position:"Query" name:"VpcId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r EnableVpcClassicLinkRequest) Invoke(client *sdk.Client) (response *EnableVpcClassicLinkResponse, err error) {
	req := struct {
		*requests.RpcRequest
		EnableVpcClassicLinkRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "EnableVpcClassicLink", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		EnableVpcClassicLinkResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.EnableVpcClassicLinkResponse

	err = client.DoAction(&req, &resp)
	return
}

type EnableVpcClassicLinkResponse struct {
	RequestId string
}
