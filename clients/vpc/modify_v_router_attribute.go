package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyVRouterAttributeRequest struct {
	requests.RpcRequest
	VRouterName          string `position:"Query" name:"VRouterName"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	VRouterId            string `position:"Query" name:"VRouterId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifyVRouterAttributeRequest) Invoke(client *sdk.Client) (resp *ModifyVRouterAttributeResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "ModifyVRouterAttribute", "vpc", "")
	resp = &ModifyVRouterAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyVRouterAttributeResponse struct {
	responses.BaseResponse
	RequestId string
}
